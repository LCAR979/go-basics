package main

import (
	"fmt"
	"math"
	"time"
)

func forLoopsBasic() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func forLoopsAdvanced() {
	sum := 1
	// or:
	/*
		for ; sum < 10; {
		...
		}
	*/
	for sum < 10 { // equal to while
		sum += sum
	}
	fmt.Println(sum)

}

func ifBasic(n int) {
	if n > 0 {
		fmt.Println("Positive")
	} else {
		fmt.Println("Negative")
	}
}

func ifAdvanced(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
	// local variables defined in `if` statement is only visible in `if` and `else` blocks
}

// `break` in each case is provided automatically,
// cases can be not constants, even not numbers
func switchCaseBasic(x float64) {
	switch x {
	case 1.0:
		fmt.Println("case 1")
	case 2.0:
		fmt.Println("case 2")
	default:
		fmt.Println("default")
	}
}

func switchCaseAdvanced() {
	t := time.Now()

	//using `switch` without condition is a clean way to write lng if-then-else chains
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 18:
		fmt.Println("Afternoon")
	default:
		fmt.Println("Night")
	}
}

// a defer statement defers the execution of a function
// until the surrounding function returns
//
// a stack is used to save defered statements
func deferUsage() {
	fmt.Println("counting")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
	// output: counting, done, 4, 3, 2, 1, 0
}

func main() {
	forLoopsBasic()
	forLoopsAdvanced()

	ifBasic(5)
	ifAdvanced(2.0, 3.0, 4.0)

	switchCaseBasic(4.2)
	switchCaseAdvanced()

	deferUsage()
}
