package gee

func newTestGroup() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hi/foo", nil)
	r.addRoute("GET", "/hi/foo/:bar", nil)
	r.addRoute("GET", "/hi/foo/:bar/*zip", nil)
	r.addRoute("GET", "/hi/bar/*zip/*", nil)
	return r
}
