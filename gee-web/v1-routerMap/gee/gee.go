package gee

import (
	"fmt"
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandleFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	path := r.URL.Path
	fullKey := method + "#" + path
	if handleFunc, ok := e.router[fullKey]; ok {
		handleFunc(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "ERROR: 404 Not Found")
	}
}

func (e *Engine) addRoute(method string, path string, handle HandleFunc) {
	fullKey := method + "#" + path
	e.router[fullKey] = handle
}

func (e *Engine) Get(path string, handle HandleFunc) {
	e.addRoute("GET", path, handle)
}

func (e *Engine) Post(path string, handle HandleFunc) {
	e.addRoute("POST", path, handle)
}

func (e *Engine) Run(addr string) (err error) {
	fmt.Println("Http service started...")
	return http.ListenAndServe(addr, e)
}
