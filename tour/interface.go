package main

import (
	"fmt"
	"math"
)

/*
Interface: a set of method signatures

A value of interface type can hold any value that implements those methods

A type implements an interface just by implmenting its methods,
no keyword or other explict declarations

underlaying structure: an interface can be thought as (val, type)
*/
type Abser interface {
	Abs() float64
}

// types that implements the Abser interface:

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyString struct {
	s string
}

func (s *MyString) Abs() float64 {
	if s == nil {
		fmt.Println("<nil>")
		return 0
	}
	fmt.Println(string(s.s))
	return 0
}

func describe(v Abser) {
	fmt.Printf("(%v, %T)\n", v, v)
}

func describeGeneral(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func basicInterfaceExample() {
	var x Abser
	x = MyFloat(-2)
	fmt.Println(x.Abs())
	describe(x) //	(-2, main.MyFloat)

	v := Vertex{X: 1, Y: 2}
	x = &v
	fmt.Println(x.Abs())
	describe(x) // (&{1 2}, *main.Vertex)
	// Note only `*vertex` type implments Abser interface, not `vertex` type
	// If use `x = v`, compile error will be thrown
}

func nilValInterfaceExample() {
	var x Abser
	var mys *MyString
	x = mys // receiver can be nil, don't forget to deal with nil in methods implementation
	x.Abs()
	describe(x)

	x = &MyString{"hello"}
	x.Abs()
	describe(x)
}

func emptyInterfaceExample() {
	// empty interface: an interface type specifies zero methods
	// and may hold values of any type
	type emptyInterface interface{}
	var ep emptyInterface
	ep = 42
	describeGeneral(ep)

	ep = "world"
	describeGeneral(ep)
}

func typeAssertExample() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	// assert the interface value holds the concrete type T(string here)
	// and get the T value
	s, ok := i.(string)
	fmt.Println(s, ok)
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// the following is wrong since the assertion is false
	// f = i.(float64)
}

func switchType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("int, value is %d\n", v)
	case string:
		fmt.Printf("string, content is %s\n", v)
	default:
		fmt.Printf("unknown type, %T\n", v)
	}
}

func switchTypeExample() {
	switchType(43)
	switchType("hello world")
	switchType(false)
	/*
		int, value is 43
		string, content is hello world
		unknown type, bool
	*/
}

// Error interface example
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("val %v is negative\n", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		e := ErrNegativeSqrt(x)
		return -1, &e
	}
	z := x
	z_new := z - (z*z-x)/(2*z)
	for math.Abs(z_new-z) > 1e-5 {
		z = z_new
		z_new = z - (z*z-x)/(2*z)
		// fmt.Printf("z = %g\n", z)
	}
	return z_new, nil
}

func main() {

	basicInterfaceExample()

	nilValInterfaceExample()

	emptyInterfaceExample()

	typeAssertExample()

	switchTypeExample()

	// error interface example
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

	var e error
	x := ErrNegativeSqrt(-2)
	e = x  // OK
	e = &x // also OK
	fmt.Println(e)
	// this related to the concept of "method set"
}

/* Further details

Method sets

A type may have a method set associated with it.
The method set of an interface type is its interface.
The method set of any other type T consists of all methods declared with receiver type T.
The method set of the corresponding pointer type *T is the set of all methods that
declared with receiver *T or T (that is, it also contains the method set of T).

Read:
https://stackoverflow.com/questions/63995894/why-can-we-assign-a-struct-pointer-to-an-interface-variable-even-though-the-stru
https://stackoverflow.com/questions/63498619/why-is-an-interface-assignment-more-strict-than-a-method-call
*/
