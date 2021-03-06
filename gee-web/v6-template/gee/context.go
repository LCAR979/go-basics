package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JSONObj map[string]interface{}

type Context struct {
	Req    *http.Request
	Writer http.ResponseWriter

	Method string
	Path   string

	// save route params
	Params map[string]string

	StatusCode int

	// middleware handlers and execute index
	handlers []HandlerFunc
	index    int

	engine *Engine
}

func newContext(req *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		Req:    req,
		Writer: w,
		Method: req.Method,
		Path:   req.URL.Path,
		index:  -1,
	}
}

func (c *Context) Next() {
	c.index++
	for ; c.index < len(c.handlers); c.index++ {
		c.handlers[c.index](c)
	}
}

//`Value` means retrieving the first value
// another version: PostFormValue(), which excludes the url query part
func (c *Context) FullFormValue(key string) string {
	return c.Req.FormValue(key)
}

// url.Query() -> Values(type like map[string][]string)
// Values.Get() returns first value in string format
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Data(code, []byte(fmt.Sprintf(format, values...)))
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Data(code, []byte(html))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) HTMLUsingTemplate(code int, name string, data interface{}) {
	c.SetHeader("Content-Type", "text/html")
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Fail(500, err.Error())
	} else {
		c.Status(code)
	}
}

func (c *Context) Param(key string) string {
	val := c.Params[key]
	return val
}

func (c *Context) Fail(code int, errMsg string) {
	c.index = len(c.handlers)
	c.JSON(code, JSONObj{"msg": errMsg})
}
