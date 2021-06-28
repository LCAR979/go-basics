package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println(runtime.NumGoroutine(), i) // 1 2 3 4
	if atomic.LoadInt32(&i) <= 0 {
		return
	}
	atomic.AddInt32(&i, -1)
	go main()
	time.Sleep(100 * time.Millisecond)
}

var i int32 = 3

/*
1 3
2 2
3 1
4 0
*/
