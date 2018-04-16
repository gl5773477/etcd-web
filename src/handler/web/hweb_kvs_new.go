package web

import (
	"context"
	"fmt"
	"handler/base"
	"net/http"
)

type HWebKvsNew struct {
	base.BaseHandler
}

func (c *HWebKvsNew) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	key := r.FormValue("key")
	if key == "" {
		base.ReponseFailed(w, "缺少key参数")
		return
	}
	fmt.Println("new 新建key：", key)

	isDir := r.FormValue("isDir")
	if isDir == "" {
		base.ReponseFailed(w, "缺少isDir参数")
		return
	}

	if isDir == "yes" {
		key += "/"
	}

	base.EtcdCli.Put(context.TODO(), key, "")

	base.ReponseSuccuss(w, nil)
}
