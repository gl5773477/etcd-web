package user

import (
	"handler/base"
	"html/template"
	"log"
	"net/http"
)

type HUserLogin struct {
	base.BaseHandler
}

func (c *HUserLogin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// if err := c.Prepare(r); err != nil {
	// 	c.SetError(503, fmt.Sprintf("err:%v", err))
	// }
	// c.Finish(w, r)

	t, err := template.ParseFiles("./conf/web/index.html")
	if err != nil {
		log.Fatal(err)
	}

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}
