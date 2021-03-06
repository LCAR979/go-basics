package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		val := a
		a, b = b, a+b
		return val
	}
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
	
	f := fibonacci()
	for i:=0; i<10; i++{
		fmt.Println(f())
	}
}
