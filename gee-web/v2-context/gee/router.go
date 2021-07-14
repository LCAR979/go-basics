package gee

import (
	"fmt"
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

type Router struct {
	table map[string]HandleFunc
}

func New() *Router {
	return &Router{table: make(map[string]HandleFunc)}
}

func (e *Router) handle(c *Context) {
	method := r.Method
	path := r.URL.Path
	fullKey := method + "#" + path
	if handleFunc, ok := e.table[fullKey]; ok {
		handleFunc(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "ERROR: 404 Not Found")
	}
}

func (r *Router) addRoute(method string, path string, handle HandleFunc) {
	fullKey := method + "#" + path
	r.table[fullKey] = handle
}
