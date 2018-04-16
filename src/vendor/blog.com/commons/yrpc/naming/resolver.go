package naming

import (
	"context"
	"encoding/json"
	"time"

	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/naming"
)

type YRPCResolver struct {
	serviceName string
	Client      *clientv3.Client
}

func NewResolver(servName string, cli *clientv3.Client) *YRPCResolver {
	return &YRPCResolver{
		serviceName: servName,
		Client:      cli,
	}
}

func (c *YRPCResolver) Regist(ctx context.Context, host string, opts ...clientv3.OpOption) (err error) {
	info := &naming.Update{
		Op:       naming.Add,
		Addr:     host,
		Metadata: c.meta(),
	}

	var bts []byte
	if bts, err = json.Marshal(info); err != nil {
		return grpc.Errorf(codes.InvalidArgument, err.Error())
	}
	_, err = c.Client.KV.Put(ctx, c.serviceName+"/"+host, string(bts), opts...)
	return
}

func (c *YRPCResolver) Delete(ctx context.Context, host string, opts ...clientv3.OpOption) error {
	_, err := c.Client.KV.Delete(ctx, c.serviceName+"/"+host, opts...)
	return err
}

func (c *YRPCResolver) Resolve(target string) (naming.Watcher, error) {
	ctx, cancel := context.WithCancel(context.Background())
	w := &watcher{
		serviceName: c.serviceName,
		cli:         c.Client,
		ctx:         ctx,
		cancel:      cancel,
	}
	return w, nil
}

// Can be used for load balance.
func (c *YRPCResolver) meta() interface{} {
	md := map[string]interface{}{
		"create_time": time.Now().Format("2006-01-02 15:04:05"),
		// TODO
	}
	bts, err := json.Marshal(md)
	if err != nil {
		return nil
	}
	return string(bts)
}
