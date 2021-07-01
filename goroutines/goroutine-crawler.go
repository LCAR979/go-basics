package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	mu sync.Mutex
	m  map[string]bool
}

func (c *ConcurrentMap) Set(s string) {
	c.mu.Lock()
	c.m[s] = true
	c.mu.Unlock()
}

func (c *ConcurrentMap) Get(s string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, exists := c.m[s]
	return exists
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, c *ConcurrentMap, wg *sync.WaitGroup) {
	// fmt.Printf("enter visiting %s \n", url)
	defer wg.Done()
	c.Set(url)

	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		visited := c.Get(u)
		if !visited {
			wg.Add(1)
			go Crawl(u, depth-1, fetcher, c, wg)
		}
	}
	return
}

func main() {
	conMap := &ConcurrentMap{m: make(map[string]bool)}
	wg := &sync.WaitGroup{}

	wg.Add(1) // this line was missed at first, leading to uncompleted execution
	go Crawl("https://golang.org/", 4, fetcher, conMap, wg)
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

/*
output:
found: https://golang.org/ "The Go Programming Language"
not found: https://golang.org/cmd/
found: https://golang.org/pkg/ "Packages"
found: https://golang.org/pkg/os/ "Package os"
found: https://golang.org/pkg/fmt/ "Package fmt"
*/
