package main

import (
	"config"
	"fmt"
	"handler"
	"handler/base"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/coreos/etcd/clientv3"
)

var (
	cli *clientv3.Client
)

func main() {
	config.Init()
	base.InitEtcd()
	fmt.Println("config info:", config.C)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	resourcePath := filepath.Join(wd, "conf/web")

	http.Handle("/", http.FileServer(http.Dir(resourcePath)))
	http.Handle("/dcmp/v1/key/list", handler.NewHWebKvsList())
	http.Handle("/dcmp/v1/key/new", handler.NewHWebKvsNew())
	http.Handle("/dcmp/v1/key/delete", handler.NewHWebKvsDelete())
	http.Handle("/dcmp/v1/key/save", handler.NewHWebKvsSave())
	http.HandleFunc("/login", index)
	http.HandleFunc("/dcmp/v1/key/export", index)
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./conf/web/index.html")
	if err := t.Execute(w, nil); err != nil {
		log.Println("template execute failed:", err)
	}
}
