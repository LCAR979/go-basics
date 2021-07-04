package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter = 0

func add(a, b int, mu *sync.Mutex) {
	c := a + b
	mu.Lock()
	counter++
	fmt.Printf("%v: %v + %v = %v\n", counter, a, b, c)
	mu.Unlock()
}

func main() {
	mu := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go add(1, i, mu)
	}

	for {
		mu.Lock()
		val := counter
		mu.Unlock()

		// runtime.Gosched() explicitly yields the control to other goroutines
		runtime.Gosched()

		if val >= 10 {
			break
		}
	}
}
