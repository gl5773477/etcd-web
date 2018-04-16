package naming

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/naming"
)

var ErrWatcherClosed = fmt.Errorf("naming: watch closed")

type watcher struct {
	serviceName string
	cli         *clientv3.Client
	ctx         context.Context
	cancel      context.CancelFunc
	wch         clientv3.WatchChan
	err         error
}

func (c *watcher) Close() { c.cancel() }
func (c *watcher) Next() ([]*naming.Update, error) {
	if c.wch == nil {
		return c.firstNext()
	}
	if c.err != nil {
		return nil, c.err
	}

	wr, ok := <-c.wch
	if !ok {
		c.err = grpc.Errorf(codes.Unavailable, "%s", ErrWatcherClosed)
		return nil, c.err
	}
	if c.err = wr.Err(); c.err != nil {
		return nil, c.err
	}

	updates := make([]*naming.Update, 0, len(wr.Events))
	for _, e := range wr.Events {
		var jupdate naming.Update
		var err error
		switch e.Type {
		case clientv3.EventTypePut:
			err = json.Unmarshal(e.Kv.Value, &jupdate)
			jupdate.Op = naming.Add
		case clientv3.EventTypeDelete:
			err = json.Unmarshal(e.PrevKv.Value, &jupdate)
			jupdate.Op = naming.Delete
		}
		if err == nil {
			updates = append(updates, &jupdate)
		}
	}
	return updates, nil
}

func (c *watcher) firstNext() ([]*naming.Update, error) {
	resp, err := c.cli.Get(c.ctx, c.serviceName, clientv3.WithPrefix())
	if c.err = err; err != nil {
		return nil, err
	}

	updates := make([]*naming.Update, 0, len(resp.Kvs))
	for _, kv := range resp.Kvs {
		var jupdate naming.Update
		if err := json.Unmarshal(kv.Value, &jupdate); err != nil {
			continue
		}
		updates = append(updates, &jupdate)
	}

	opts := []clientv3.OpOption{
		clientv3.WithRev(resp.Header.Revision + 1),
		clientv3.WithPrefix(),
		clientv3.WithPrevKV(),
	}

	c.wch = c.cli.Watch(c.ctx, c.serviceName, opts...)
	return updates, nil
}
