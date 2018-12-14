package ey

import (
	"container/list"
	"io"
	"net/http"
)

type Ctx interface {
	Next()
	Out(string)
	io.Writer
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
func (ctx *Context) Write(buf []byte) (int, error) {
	return ctx.W.Write(buf)
}
