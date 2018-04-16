package base

import (
	"fmt"
	"net/http"

	"blog.com/render"
)

type Result struct {
	Data    interface{} `json:"data"`
	RetCode int32       `json:"retCode"`
	RetMsg  string      `json:"retMsg"`
}

func ReponseFailed(w http.ResponseWriter, msg string) {
	res := &Result{}
	res.RetCode = 500
	res.RetMsg = msg
	js := render.JSON{
		Data: res,
	}
	err := js.Render(w)
	if err != nil {
		fmt.Println("[测] 渲染失败")
		return
	}
}

func ReponseSuccuss(w http.ResponseWriter, resp interface{}) {
	res := &Result{}
	res.RetCode = 0
	res.RetMsg = ""
	res.Data = resp
	js := render.JSON{
		Data: res,
	}
	err := js.Render(w)
	if err != nil {
		fmt.Println("[测] 渲染失败")
	}
}

func ReponseSuccussWithHint(w http.ResponseWriter, code int32, msg string, resp interface{}) {
	res := &Result{}
	res.RetCode = 0
	res.RetMsg = ""
	res.Data = resp
	js := render.JSON{
		Data: res,
	}
	err := js.Render(w)
	if err != nil {
		fmt.Println("[测] 渲染失败")
	}
}
