package yrpc

import (
	"github.com/coreos/etcd/clientv3"
	// "golang.org/x/net/trace"
)

func NewEtcdConn(cfg *EtcdConfig) *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: cfg.DialTimeout,
	})
	if err != nil {
		panic(err)
	}
	return cli
}
