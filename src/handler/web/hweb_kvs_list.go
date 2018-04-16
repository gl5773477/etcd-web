package web

import (
	"context"
	"fmt"
	"handler/base"
	"net/http"
	"strings"

	"github.com/coreos/etcd/client"
	"github.com/coreos/etcd/clientv3"
)

type HWebKvsList struct {
	base.BaseHandler
}

func (c *HWebKvsList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	treeNode := formatEtcdNodes(nil, true)

	base.ReponseSuccuss(w, treeNode)
}

func formatEtcdNodes(node *client.Node, isDir bool) (treeNode *TreeNode) {
	resp, err := base.EtcdCli.Get(context.TODO(), "config", clientv3.WithPrefix())
	if err != nil {
		fmt.Printf("获取key失败：%v\n", err)
		return
	}

	var keyMp = make(map[string][]string)
	var valueMp = make(map[string]string)
	for _, kv := range resp.Kvs {
		// fmt.Println("[测] k=", string(kv.Key), "v=", string(kv.Value))
		k := string(kv.Key)
		sli := strings.Split(k, "/")
		for _, v := range sli {
			keyMp[k] = append(keyMp[k], v)
		}
		valueMp[k] = string(kv.Value)
	}

	treeNode = NewNode("config", "config", "")
	for key, idxList := range keyMp {
		treeNode.AddChild(idxList, valueMp[key])
		fmt.Println("目录节点树：", treeNode.ToString())
	}
	return
}
