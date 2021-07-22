package gee

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hi/foo", nil)
	r.addRoute("GET", "/hi/foo/:bar", nil)
	r.addRoute("GET", "/hi/foo/:bar/*zip", nil)
	r.addRoute("GET", "/hi/bar/*zip/*", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok1 := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok2 := reflect.DeepEqual(parsePattern("/p/*path"), []string{"p", "*path"})
	ok3 := reflect.DeepEqual(parsePattern("/p/*path/*"), []string{"p", "*path"})
	ok4 := reflect.DeepEqual(parsePattern("/p/foo/bar/zip"), []string{"p", "foo", "bar", "zip"})
	fmt.Println(ok1, ok2, ok3, ok4)
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, params := r.getRoute("GET", "/hi/foo/x/y")
	if n == nil {
		t.Fatal("node matching the pattern is not found")
	}
	if n.totalPattern != "/hi/foo/:bar/*zip" {
		fmt.Println(n)
		t.Fatal("pattern saved in leaf node is wrong")
	}
	if params["bar"] != "x" || params["zip"] != "y" {
		t.Fatal("params evaluated error")
	}

	fmt.Println("testing getRoute() all passed")
}
