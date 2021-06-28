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

func basicExample() {
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
	// go sumOnSlice(arr[:len(arr)/2], c)
	// sumOnSlice(arr[len(arr)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x + y)
	fmt.Println(arr)
}

func main() {
	basicExample()
	channelExample()
}

/* Further Reading

GO programs ends when the main function ends


https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb

*/
