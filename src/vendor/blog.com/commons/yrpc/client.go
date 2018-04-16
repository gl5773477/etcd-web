package yrpc

import (
	"context"
	"fmt"
	"time"

	"git.ymt360.com/go/gocommons/yrpc/naming"
	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc"
)

type YRPCClient struct {
	cli *clientv3.Client
}

func NewClient(cfg *EtcdConfig) *YRPCClient {
	return &YRPCClient{
		cli: NewEtcdConn(cfg),
	}
}

func (c *YRPCClient) Connect(servName string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	r := naming.NewResolver(servName, c.cli)
	b := RoundRobin(r)

	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBalancer(b))
	opts = append(opts, grpc.WithUnaryInterceptor(ClientInterceptor))

	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	conn, err := grpc.DialContext(ctx, servName, opts...)
	if err != nil {
		return nil, fmt.Errorf("did not connect %s: %v", servName, err)
	}

	return conn, nil
}
