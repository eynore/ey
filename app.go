package ermao

import (
	"container/list"
	"net/http"
)

type Path string

type Handler func(Ctx)

type PathMap map[Path][]Handler
type MethodMap map[string]PathMap
type App struct {
	methodMap MethodMap
}

func New() *App {
	return &App{make(MethodMap)}
}

func (app *App) reg(method string, path Path, handlers ...Handler) {
	pathMap := app.methodMap[method]
	if pathMap == nil {
		app.methodMap[method] = make(PathMap)
		pathMap = app.methodMap[method]
	}
	pathMap[path] = handlers
}

func (app *App) Get(path Path, handlers ...Handler) {
	app.reg("GET", path, handlers...)
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathMap := app.methodMap[r.Method]
	if pathMap == nil {
		return
	}
	handlers := pathMap[Path(r.URL.Path)]

	if len(handlers) == 0 {
		return
	}

	list := list.New()
	for _, fn := range handlers {
		list.PushBack(fn)
	}
	fn := list.Front()
	var ctx *Context
	ctx = &Context{w, r, fn}
	fn.Value.(Handler)(ctx)
}
func (app *App) Listen(addr string) {
	http.ListenAndServe(addr, app)
}
