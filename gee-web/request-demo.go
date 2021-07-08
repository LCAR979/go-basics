package main

import (
	"fmt"
	"log"
	"net/url"
)

/*https://pkg.go.dev/command-line-arguments?utm_source=gopls#UrlQueryUsage
func (u *URL) Query() Values
Query parses RawQuery and returns the corresponding values.
It silently discards malformed value pairs. To check errors use ParseQuery.
*/

func UrlQueryUsage() {
	u, err := url.Parse("https://example.org/?a=1&a=2&b=&=3&&&&")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	fmt.Println(q)
	fmt.Println(q["a"])
	fmt.Println(q.Get("a"))
	fmt.Println(q.Get("b"))
	fmt.Println(q.Get(""))
	/*result
	map[:[3] a:[1 2] b:[]]
	[1 2]
	1

	3
	*/
}

func main() {
	UrlQueryUsage()
}
