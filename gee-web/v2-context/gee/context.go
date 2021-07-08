package context

import "net/http"

type context struct {
	Req    *http.Request
	Writer http.ResponseWriter

	Method string
	Path   string

	StatusCode int
}

func newContext(req *http.Request, w http.ResponseWriter) *context {
	return &context{
		Req:    req,
		Writer: w,
		Method: req.Method,
		Path:   req.URL.Path,
	}
}

//`Value` means retrieving the first value
// another version: PostFormValue(), which excludes the url query part
func (c *context) FormValue(key string) {
	return c.Req.FormValue(key)
}

// url.Query() -> Values(type like map[string][]string)
// Values.Get() returns first value in string format
func (c *context) Query(key string) string {
	return c.Req.url.Query().Get(key)
}

func (c *context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}
