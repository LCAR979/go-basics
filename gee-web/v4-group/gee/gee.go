package gee

import (
	"fmt"
	"net/http"
)

type HandleFunc func(c *Context)

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}
type RouterGroup struct {
	prefix      string
	middlewares []HandleFunc
	parent      *RouterGroup
	engine      *Engine
}

func NewEngine() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) addRoute(method string, path string, handler HandleFunc) {
	e.router.addRoute(method, path, handler)
}

func (e *Engine) GET(path string, handle HandleFunc) {
	e.addRoute("GET", path, handle)
}

func (e *Engine) POST(path string, handle HandleFunc) {
	e.addRoute("POST", path, handle)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(req, w)
	e.router.handle(c)
}

func (e *Engine) Run(addr string) (err error) {
	fmt.Println("Http service started...")
	return http.ListenAndServe(addr, e)
}
