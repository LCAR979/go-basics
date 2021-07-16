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
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.JSONObj{
			"username": c.FullFormValue("username"),
			"password": c.FullFormValue("password"),
		})
	})

	r.Run(":9999")
}
