package gee

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{roots: make(map[string]*node),
		handlers: make(map[string]HandlerFunc)}
}

// parse a path pattern like `/hello/:name`
// which: splits each part and save in a slice, will drop parts after the first *
func parsePattern(pattern string) []string {
	splits := strings.Split(pattern, "/")
	res := make([]string, 0)
	for _, s := range splits {
		if s != "" {
			res = append(res, s)
			if s[0] == '*' {
				break
			}
		}
	}
	return res
}

// addRoute() should be called by developers to build the whole router trie
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = new(node)
	}

	parts := parsePattern(pattern)
	r.roots[method].insert(pattern, parts, 0)

	key := method + "#" + pattern
	r.handlers[key] = handler
}

// getRoute() is used when a new visit from client comes,
// and path is parsed to match patterns saved in trie
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	params := make(map[string]string)
	searchParts := parsePattern(path)

	n := root.search(searchParts, 0)
	if n == nil {
		return nil, nil
	}

	nodeParts := parsePattern(n.totalPattern)
	for idx, part := range nodeParts {
		if part[0] == ':' {
			params[part[1:]] = searchParts[idx]
		}
		if part[0] == '*' && len(part) > 1 {
			params[part[1:]] = strings.Join(searchParts[idx:], "/")
			break
		}
	}
	return n, params
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		// don't forget to save params
		c.Params = params
		handleFunc := r.handlers[c.Method+"#"+n.totalPattern]
		handleFunc(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND")
	}
}
