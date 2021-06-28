package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.NumGoroutine()) // 3
	time.Sleep(100 * time.Millisecond)
}
func init() {
	go main()
	go main()
}

/*
3
3
3

https://stackoverflow.com/questions/53388154/is-the-main-function-runs-as-a-goroutine
*/
