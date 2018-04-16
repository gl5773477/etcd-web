package base

import (
	"fmt"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
)

var (
	EtcdCli *clientv3.Client
)

func init() {
	var err error
	EtcdCli, err = clientv3.New(clientv3.Config{
		Endpoints:   strings.Split("10.10.91.206:12379", ","),
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Println("连接etcd失败：", err)
		return
	}
}
