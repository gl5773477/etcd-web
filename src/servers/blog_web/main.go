package main

import (
	"fmt"
	"handler"
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
	// http.Handle("/dcmp/v1/key/export", handler.NewHWebKvsExport())
	http.HandleFunc("/login", index)
	http.HandleFunc("/dcmp/v1/key/export", index)
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./conf/web/index.html")
		if err := t.Execute(w, nil); err != nil {
			fmt.Println("[测] ======template failed:", err)
		}

	} else {
		t, _ := template.ParseFiles("./conf/web/index.html")
		if err := t.Execute(w, nil); err != nil {
			fmt.Println("[测] ======template failed:", err)
		}

	}
}
