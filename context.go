package ey

import (
	"container/list"
	"net/http"
)

type Ctx interface {
	Next()
	Out(string)
}

type Context struct {
	W  http.ResponseWriter
	R  *http.Request
	el *list.Element
}

func (ctx *Context) Next() {
	if el := ctx.el.Next(); el != nil {
		el.Value.(func(Ctx))(ctx)
	}
}

func (ctx *Context) Out(str string) {
	ctx.W.Write([]byte(str))
}
