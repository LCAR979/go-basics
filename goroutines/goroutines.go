package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func basicGoroutineExample() {
	go say("hello")
	say("world")
}

func sumOnSlice(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func channelExample() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	c := make(chan int)
	go sumOnSlice(arr[:len(arr)/2], c)
	go sumOnSlice(arr[len(arr)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x + y)
	fmt.Println(arr)
}

func fibonacci(c chan int, n int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // only the sender should close the channel
}

func closeAndRangeUsageExample() {
	c := make(chan int, 10)
	go fibonacci(c, cap(c))
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacciQuitsOnSignal(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func selectExample() {
	c := make(chan int, 10)
	quit := make(chan int, 10)
	go fibonacciQuitsOnSignal(c, quit)
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	quit <- 0
	time.Sleep(1000 * time.Millisecond)
}

func defaultExample() {
	tick := time.Tick(100 * time.Millisecond)
	after := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-after:
			fmt.Println("time up")
			return
		default: // The default case in a select is run if no other case is ready.
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	fmt.Println("-------basic goroutine example--------")
	basicGoroutineExample()

	fmt.Println("-------channel example--------")
	channelExample()

	fmt.Println("-------close and range example--------")
	closeAndRangeUsageExample()

	fmt.Println("-------select example--------")
	selectExample()

	fmt.Println("-------default example--------")
	defaultExample()
}

/* Further Reading

GO programs ends when the main function ends


https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb go的channel详解

*/
