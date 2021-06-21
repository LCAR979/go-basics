package main

import (
	"fmt"
)

//pointer
func pointerUsage() {
	fmt.Println("\n--------------------\npointer usage")
	var p *int
	i := 42
	p = &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(*p)
}

// struct
type vertex struct {
	x int
	y int
}

func structUsage() {
	fmt.Println("\n--------------------\nstruct usage")
	fmt.Println(vertex{x: 1, y: 2})
	v := vertex{x: 1, y: 2}
	//var v vertex = vertex{x: 1, y:2}
	v.x = 2 // set value using `.` (dot)
	fmt.Println(v)

	p := &v  // p is a pointer to v
	p.x = 10 // NOTE: (*p).field is omitted to p.field
	fmt.Println(*p)
}

// array & slice
func arrayUsage() {
	fmt.Println("\n--------------------\narray usage")

	// an array's length is part if its type
	// so array cannot be resized
	a := [5]int{1, 2, 4, 6, 8} // this is called array literals

	// **********NOTE!!!!!!!!!**********
	// when you assign or pass around an array value,
	// you will make a copy of its contents.
	b := a // b is a copy to a's contents
	b[0] = 100
	fmt.Println(a, b) //different outputs

	// another writing method:
	// var b [4]int = [4]int{1, 2, 3}
	// If you want to initialize at declaration time, use :=
}

func sliceUsage() {
	fmt.Println("\n--------------------\nslice usage")
	a := [5]int{1, 2, 4, 6, 8}
	// slice expresion can use default values
	fmt.Println(a[0:5], a[0:], a[:5], a[:]) // same outputs

	var s []int = a[0:4] // include left end, exclude right end
	// using make to create a slice, syntax: make([]T, len, cap)
	// `make` allocates an array and returns a slice that refers to that array.
	// s2 := make([]int, 5, 5)

	fmt.Println(a, s, len(s), cap(s)) // len = 4, cap = 5
	// besides, using s[4] is wrong, since indexing is based on visible? range

	structSlice := []struct {
		a int
		b string
	}{
		{1, "h"},
		{2, "l"},
	}
	fmt.Println(structSlice)

	structSlice2 := []vertex{
		{1, 2},
		{3, 5},
	}
	fmt.Println(structSlice2)

}

// when you change the content of a slice,
// both the original array and other slices will be affected
func sliceLikeReferenceToArray() {
	fmt.Println("\n--------------------\nslice like reference to array")

	// slice literals
	a := []string{"hello", "world", "hello", "human"}
	var s1 = a[0:2]
	var s2 = a[1:4]
	fmt.Println(a, s1, s2)
	s1[1] = "sekai"
	fmt.Println(a, s1, s2)
}

func showSliceLenAndCap(s []int) {
	fmt.Printf("Len: %d, Cap: %d\n", len(s), cap(s))
}

func sliceLenAndCap() {
	fmt.Println("\n--------------------\nslice len and cap")
	a := [5]int{1, 2, 3, 4, 5}
	s := a[0:4]
	// when calculating cap, use prev_cap - slice_start_index
	showSliceLenAndCap(s) // s = {1, 2, 3, 4} len = 4, cap = 5, 5-0 = 5
	s = s[1:]
	showSliceLenAndCap(s) // s = {2, 3, 4},   len = 3, cap = 4, 5-1 = 4
	s = s[:len(s)-1]
	showSliceLenAndCap(s) // s = {2, 3},	  len = 2, cap = 4, 4-0 = 4
}

func loopSlice() {
	fmt.Println("\n--------------------\nloop slice")
	a := []int{1, 2, 3}
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()
	for i, val := range a {
		fmt.Printf("a[%d] = %d\n", i, val)
	}
}

func appendSlice() {
	fmt.Println("\n--------------------\nappend slice")
	a := []int{1, 2, 3}
	a = append(a, 4, 5)
	b := []int{6, 7}

	// **********NOTE!!!**********
	// when doing slice and slice append, append `...` after those params to merge
	// otherwise will cause error
	a = append(a, b...)
	fmt.Println(a, len(a), cap(a)) // len = 7, cap = 12
	// when capacity is not enough => copy to a new underlaying array, cap is doubled!
}

func copySlice() {
	fmt.Println("\n--------------------\ncopy slice")

	/*
		copy function: copy(dst, src []T) int
		return the number of elements successfully copied: min(len(dst), len(src))
		once copied, they are using two different underlaying arrays
	*/

	var a, b []int
	var n int
	a = make([]int, 4)
	b = []int{5, 6}
	n = copy(a, b)
	a[0] = 100
	fmt.Println(a, b, n)
	// [100 6 0 0] [5 6] 2

	a = make([]int, 2)
	b = []int{5, 6, 7, 8}
	n = copy(a, b)
	a[0] = 100
	fmt.Println(a, b, n)
	// [100, 6] [5 6 7 8] 2
}

func twoDSliceUsage() {
	fmt.Println("\n--------------------\ntwo slice usage")
	twoDSlice := make([][]int, 2)

	// you have to specify inner dimension
	for i := range twoDSlice {
		twoDSlice[i] = make([]int, 3)
	}

	// or: using slice literals 
	// NOTE: inner slices can have different dimensions
	twoDSlice[0] = []int{1, 2, 3}
	twoDSlice[1] = []int{4, 6, 7, 9}
	fmt.Println(twoDSlice)
	fmt.Printf("row: %d, col-0: %d, col-1: %d\n", 
	len(twoDSlice), len(twoDSlice[0]), len(twoDSlice[1]))
}

func tutorialExample() {
	fmt.Println("\n--------------------\ntutorial example")
	s := []int {2, 3, 5, 7, 11, 13}
	showSliceLenAndCap(s)

	s = s[:0]
	showSliceLenAndCap(s)

	s = s[:4]
	showSliceLenAndCap(s)

	s = s[2:]
	showSliceLenAndCap(s)
}

func main() {
	pointerUsage()
	structUsage()
	arrayUsage()
	sliceUsage()
	sliceLikeReferenceToArray()
	sliceLenAndCap()
	loopSlice()
	appendSlice()
	copySlice()
	twoDSliceUsage()
	tutorialExample()
}

/* --------------------
Read More:
https://golangbyexample.com/slice-in-golang/
-------------------------
*/
