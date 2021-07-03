package main

import (
	"fmt"
	"log"
	"net/http"

	"gee"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.path = %q\n", r.URL.Path)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Println(k, v)
		fmt.Fprintf(w, "Header[%q] = %v\n", k, v)
	}
}

func main() {
	engine := gee.New()
	engine.Get("/", indexHandler)
	engine.Get("/hello", helloHandler)
	log.Fatal(engine.Run(":8085"))
}
