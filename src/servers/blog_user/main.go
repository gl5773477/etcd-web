package main

import (
	"fmt"
	"handler"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// sqlCfg := &config.SQLConfig{}

	// db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", sqlCfg.User, sqlCfg.Password, sqlCfg.Host, sqlCfg.Port, sqlCfg.Db))
	// if err != nil {
	// 	log.Fatal("init DB error:", err)
	// }
	// db.LogMode(true)
	// db.DB().SetMaxIdleConns(1 << 3)

	http.Handle("/login", handler.NewHUser())
	if err := http.ListenAndServe(":56001", nil); err != nil {
		fmt.Printf("server listen failed:%v\n", err)
	}
}
