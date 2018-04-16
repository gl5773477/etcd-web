package handler

import (
	"handler/web"
)

func NewHWebKvsList() *web.HWebKvsList {
	return &web.HWebKvsList{}
}

func NewHWebKvsNew() *web.HWebKvsNew {
	return &web.HWebKvsNew{}
}

func NewHWebKvsDelete() *web.HWebKvsDelete {
	return &web.HWebKvsDelete{}
}

func NewHWebKvsSave() *web.HWebKvsSave {
	return &web.HWebKvsSave{}
}

func NewHWebKvsExport() *web.HWebKvsExport {
	return &web.HWebKvsExport{}
}
