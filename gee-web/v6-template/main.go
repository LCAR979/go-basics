package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"gee"
)

type student struct {
	Name string
	Age  int8
}

func FormatDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

// developer usage
func main() {
	r := gee.NewEngine()
	r.AddMiddleware(gee.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	v1 := r.NewGroup("/v1")
	{
		v1.GET("/hello", func(c *gee.Context) {
			// expect /hello?name=arthur
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
		v1.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/arthur
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	v2 := r.NewGroup("/v2")
	v2.AddMiddleware(gee.LoggerForV2())
	{
		v2.GET("/assets/*filepath", func(c *gee.Context) {
			// expect /assets/css/index.css
			c.JSON(http.StatusOK, gee.JSONObj{"filepath": c.Param("filepath")})
		})

		authNode := v2.NewGroup("/auth")
		authNode.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.JSONObj{
				"username": c.FullFormValue("username"),
				"password": c.FullFormValue("password"),
			})
		})
	}

	r.Run(":9999")
}
