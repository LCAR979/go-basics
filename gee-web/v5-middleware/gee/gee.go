package gee

import (
	"fmt"
	"log"
	"net/http"
)

type HandleFunc func(c *Context)

/* the reason of `engine` field here:
we want to give `routerGroup` the ability to create route policies,
by adding the `engine` field, we can use routerGroup.engine.router.addRoute to achieve the goal.
And we also want to keep the ability to directly control a route policy not related to one certain routerGroup,
so we cannot just move the `router` to the inner of `RouterGroup`
*/
type RouterGroup struct {
	prefix      string
	middlewares []HandleFunc
	engine      *Engine // pointing to the only one global engine istance
}

/*Embedded struct
type Engine struct {
	rg *RouterGroup
	router *Router
	groups []*RouterGroup
}
func (engine *Engine) GET(path string, handler HandlerFunc) {
	engine.rg.addRoute("GET", pattern, handler)
}
*/

type Engine struct {
	*RouterGroup // embeded type
	router       *router
	groups       []*RouterGroup
}

func NewEngine() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) NewGroup(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, urlPattern string, handler HandleFunc) {
	pattern := group.prefix + urlPattern
	log.Printf("pattern = %s\n", pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandleFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandleFunc) {
	group.addRoute("POST", pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(req, w)
	e.router.handle(c)
}

func (e *Engine) Run(addr string) (err error) {
	fmt.Println("Http service started...")
	return http.ListenAndServe(addr, e)
}
