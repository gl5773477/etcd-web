package web

import (
	"encoding/json"
	"fmt"
)

// 允许存在重复key的节点树
type TreeNode struct {
	Key     string      `json:"name"`
	Value   string      `json:"value"`
	Dir     bool        `json:"dir"`
	Toggled bool        `json:"toggled"`
	Path    string      `json:"path"`
	Nds     []*TreeNode `json:"children"`
}

func NewNode(path, k, v string) *TreeNode {
	nd := new(TreeNode)
	nd.Key = k
	nd.Dir = true
	nd.Path = path
	nd.Value = v
	nd.Toggled = true
	nd.Nds = make([]*TreeNode, 0)
	return nd
}

func (c *TreeNode) AddChild(keys []string, value string) *TreeNode {
	parent, path := c.CreateTreeDir(keys)

	k := keys[len(keys)-1]
	if child := parent.HorizonSeach(k); child != nil { // 已存在该叶子节点
		fmt.Println("已存在叶子节点：", k)
		return nil
	}

	path += k
	var leafNode *TreeNode
	if k != "" { // 叶子
		leafNode = NewNode(path, k, value)
		leafNode.Nds = nil
		leafNode.Dir = false
		parent.Nds = append(parent.Nds, leafNode)
	} else { // 目录
		parent.Value = value
		parent.Dir = true
	}
	return leafNode
}

// 创建目录树，返回末端目录节点和树链路
func (c *TreeNode) CreateTreeDir(keys []string) (parent *TreeNode, path string) {
	height := len(keys)

	parent = c
	for i := 0; i < height-1; i++ {
		path += keys[i]
		child := parent.HorizonSeach(keys[i])
		if child != nil {
			parent = child
		} else { // 添加一个中间节点
			if keys[i] != c.Key { // 根节点不重复创建
				nd := NewNode(path, keys[i], "")
				parent.Nds = append(parent.Nds, nd)
				parent = nd
			}
		}
		path += "/"
	}
	return
}

// HorizonSeach 水平查找
func (c *TreeNode) HorizonSeach(key string) *TreeNode {
	for _, child := range c.Nds {
		if child.Key == key {
			return child
		}
	}
	return nil
}

// ToString 字符串形式串联树结构：key_path_dir
func (c *TreeNode) ToString() string {
	mp := c.dirInfos()
	bts, _ := json.Marshal(mp)
	return string(bts)
}

func (c *TreeNode) dirInfos() map[string]interface{} {
	var mp = make(map[string]interface{})
	for _, nd := range c.Nds {
		mp[nd.Key+"_["+nd.Path+"]_"+fmt.Sprintf("%t", nd.Dir)] = nd.dirInfos()
	}
	return mp
}
