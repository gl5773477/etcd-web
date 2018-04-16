package base

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BaseHandler struct {
	Req  interface{}
	Resp interface{}
}

func (c *BaseHandler) Prepare(r *http.Request) error {
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	fmt.Println("【Request】data:", string(req))
	c.Req = req
	return nil
}

func (c *BaseHandler) Finish(w http.ResponseWriter, r *http.Request) error {
	if c.Resp == nil {
		c.Resp = map[string]interface{}{
			"status": 0,
			"msg":    "",
			"data":   "",
		}
	}

	js, err := json.Marshal(c.Resp)
	if err != nil {
		fmt.Printf("json marshal failed:%v", err)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return nil
}

func (c *BaseHandler) SetError(errCode int, msg string) {
	resp := map[string]interface{}{
		"status": errCode,
		"msg":    msg,
	}

	js, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("json marshal failed:%v", err)
		return
	}

	c.Resp = js
}
