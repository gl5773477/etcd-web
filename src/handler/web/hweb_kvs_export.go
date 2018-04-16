package web

import (
	"handler/base"
	"net/http"
)

type HWebKvsExport struct {
	base.BaseHandler
}

func (c *HWebKvsExport) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	treeNode := formatEtcdNodes(nil, true)

	base.ReponseSuccuss(w, treeNode)
}
