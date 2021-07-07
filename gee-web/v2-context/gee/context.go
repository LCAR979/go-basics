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

func (c *context) PostFormValue(key string) {

}

func (c *context) Query(key string) {

}

func (c *context) Status(code int) {

}
