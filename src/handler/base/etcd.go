package base

import (
	"config"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

var (
	EtcdCli *clientv3.Client
)

func InitEtcd() {
	var err error
	EtcdCli, err = clientv3.New(clientv3.Config{
		Endpoints:   config.C.Endpoints,
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Println("连接etcd失败：", err)
		return
	}
}
