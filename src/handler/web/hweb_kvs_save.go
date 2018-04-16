package web

import (
	"context"
	"fmt"
	"handler/base"
	"net/http"

	"github.com/coreos/etcd/clientv3"
)

type HWebKvsSave struct {
	base.BaseHandler
}

func (c *HWebKvsSave) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	if key == "" {
		base.ReponseFailed(w, "缺少key参数")
		return
	}

	value := r.FormValue("value")
	if value == "" {
		base.ReponseFailed(w, "缺少value参数")
		return
	}

	fmt.Println("save 更新key：", key)
	putResp, _ := base.EtcdCli.Put(context.TODO(), key, value, clientv3.WithPrevKV())

	type ret struct {
		Value string
	}

	base.ReponseSuccussWithHint(w, 200, "保存成功", &ret{
		Value: string(putResp.PrevKv.Value),
	})
}
