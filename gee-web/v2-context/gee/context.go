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

	StatusCode int
}

func newContext(req *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		Req:    req,
		Writer: w,
		Method: req.Method,
		Path:   req.URL.Path,
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
		panic(err)
	}
}
