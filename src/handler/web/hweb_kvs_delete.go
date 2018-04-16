package web

import (
	"context"
	"fmt"
	"handler/base"
	"net/http"

	"github.com/coreos/etcd/client"
	"github.com/coreos/etcd/clientv3"
)

type HWebKvsDelete struct {
	base.BaseHandler
}

func (c *HWebKvsDelete) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	if key == "" {
		base.ReponseFailed(w, "缺少key参数")
		return
	}

	fmt.Println("delete 删除key：", key)
	base.EtcdCli.Delete(context.TODO(), key, clientv3.WithPrefix())

	base.ReponseSuccussWithHint(w, 200, "删除成功", &client.Response{})
}
