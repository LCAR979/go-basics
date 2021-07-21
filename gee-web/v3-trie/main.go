package main

import (
	"net/http"

	"gee"
)

func main() {
	r := gee.NewEngine()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=arthur
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.GET("/hello/:name", func(c *gee.Context) {
		// expect /hello/arthur
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gee.Context) {
		// expect /assets/css/index.css
		c.Json(http.StatusOK, gee.jsonObj{"filepath": c.Param("filepath")})
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.JSONObj{
			"username": c.FullFormValue("username"),
			"password": c.FullFormValue("password"),
		})
	})

	r.Run(":9999")
}
