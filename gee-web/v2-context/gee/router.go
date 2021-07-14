package gee

import (
	"log"
	"net/http"
)

type router struct {
	table map[string]HandleFunc
}

func newRouter() *router {
	return &router{table: make(map[string]HandleFunc)}
}

func (r *router) addRoute(method string, path string, handle HandleFunc) {
	log.Printf("Route %4s - %s\n", method, path)
	fullKey := method + "#" + path
	r.table[fullKey] = handle
}

func (r *router) handle(c *Context) {
	method := c.Method
	path := c.Path
	fullKey := method + "#" + path
	if handleFunc, ok := r.table[fullKey]; ok {
		handleFunc(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND")
	}
}
