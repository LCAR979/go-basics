package context

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


