package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func trace(msg string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:])
	var res strings.Builder
	res.WriteString(msg + " Trackback:\n")

	for _, pc := range pcs[n:] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		res.WriteString(fmt.Sprintf("\t%s: %d\n", file, line))
	}
	return res.String()
}

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				msg := fmt.Sprintf("%s", err)
				log.Printf("%s\n", trace(msg))
				c.Fail(http.StatusInternalServerError, "Internal server error")
			}
		}()
		c.Next()
	}
}
