package main

import (
	"fmt"
	"math"
)

/* Introduction
A method is a function with a special receiver argument

Reason to choose pointer receiver:
- When you want to modify the value
- When the receiver is a large struct, to save the time copying
*/
type vertex struct {
	x, y float64
}

// you can delcate a method on a non-struct type, too
type myFloat float64

func sqrt(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func (v vertex) abs() float64 {
	return sqrt(v.x, v.y)
}

func absV2(v vertex) float64 {
	return sqrt(v.x, v.y)
}

// methods example on non-struct types
func (val myFloat) floatAbs() float64 {
	if val < 0 {
		return float64(-val)
	}
	return float64(val)
}

// methods example of pointer receivers
func (v *vertex) scale() {
	v.x *= 10
	v.y *= 10
}

func scaleFunc(v *vertex) {
	v.x *= 10
	v.y *= 10
}

func main() {
	v1 := vertex{x: 1, y: 2}
	fmt.Println(v1.abs())
	fmt.Println(absV2(v1))

	val1 := myFloat(-2)
	fmt.Println(val1.floatAbs())

	/*
		methods with pointer receivers take either a value or a pointer
		as the receiver when they are called ==> both OK when you call the method

		but for function calls, you must use appropriate parameters,
		which should be of the same with the function defition
	*/
	v1.scale()
	p := &v1
	p.scale() // scale a second time
	fmt.Println(v1.abs())

	scaleFunc(&v1)
	fmt.Println(v1.abs())
}
