package gee

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
)

type HandlerFunc func(c *Context)

/* the reason of `engine` field here:
we want to give `routerGroup` the ability to create route policies,
by adding the `engine` field, we can use routerGroup.engine.router.addRoute to achieve the goal.
And we also want to keep the ability to directly control a route policy which is not related to one certain routerGroup,
so we cannot just move the `router` field to the inner of `RouterGroup`
*/
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
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

func (group *RouterGroup) AddMiddleware(middleware ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middleware...)
}

func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		fileName := c.Param("filepath")
		if _, err := fs.Open(fileName); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	group.GET(urlPattern, handler)
}

func (group *RouterGroup) addRoute(method string, urlPattern string, handler HandlerFunc) {
	pattern := group.prefix + urlPattern
	log.Printf("pattern = %s\n", pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc

	reqUrl := req.URL.Path
	for _, group := range e.groups {
		if strings.HasPrefix(reqUrl, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	c := newContext(req, w)
	c.handlers = middlewares
	e.router.handle(c)
}

func (e *Engine) Run(addr string) (err error) {
	fmt.Println("Http service started...")
	return http.ListenAndServe(addr, e)
}
