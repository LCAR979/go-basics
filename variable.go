package main

// import a package doesn't import its subpackages
import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/* data types in go
bool
string
int,  int8,  int16,  int32,  int64
uint, uint8, uint16, uint32, uint64
byte // alias for uint8
rune // alias for int32
float32, float64
complex64, complex128

note assignment between items of different types
requies an explict conversion
*/

// Outside a function, every statement begins with a keyword
// (var, func, and so on) and so the := construct is not available.
var pkgLevelVar1, pkgLevelVar2 bool
var varWithInitVal = 1
var varWithInitVal2 int = 2

// var can be grouped
var (
	VarGrouped1 int
	VarGrouped2 bool
	VarGrouped3 string
	// Note these two vars can be exported because starting with a Capatical Letter
)

const (
	constA         = 12
	constB float32 = 24 << 100
)

// show default data values
func showDefaultValue() {
	fmt.Println("default value of each type:")
	fmt.Println("int: ", VarGrouped1)
	fmt.Println("bool: ", VarGrouped2)
	fmt.Println("str: ", VarGrouped3)
}

// infer data types when not specified at declaration
func inferDataTypes() {
	v := 42
	fmt.Printf("v is of type %T when use 42\n", v)
	vv := 42.0
	fmt.Printf("vv is of type %T when use 42.0\n", vv)
}

// use const
// const cannot be declared using :=
// numeric consts are high-precision values
func useConsts() {
	fmt.Printf("Const value: %d, %f", constA, constB)
}

// for function params: name first, type after
// return value at the end
func add(x int, y int) int {
	return x + y
}

// params of same type: the type keyword can be merged,
// return any number of results
func mergeTypeReturnMultiple(x, y int, z float32, d, e int) (int, float32) {
	return int(z + float32(d) + float32(e)), float32(x + y)
}

func explictReturnType(x, y int) (a, b int) {
	return x + y, x - y
	/*
		naked return: return without arguments returns the named return values
		should be only used in small functions
		`return`
	*/
}

func main() {

	// Seed should not be called concurrently with any other Rand method.
	rand.Seed(time.Now().Unix())
	fmt.Println("number is ", rand.Intn(10))
	fmt.Println("number is ", rand.Intn(10))

	// a name is exported if it begins with a capital letter
	// and when importing a package, you can only use those exported names
	// i.e. must in the format of "importPackage.CapitalLetter"
	fmt.Println(math.Pi)

	fmt.Println("1+2 = ", add(1, 2))

	var a, b int
	var c float32

	//using := can defer its type
	k := 32

	a, c = mergeTypeReturnMultiple(1, 2, 3, 4, 5)
	fmt.Println("merge type of params, return multiple values: ", a, c)

	a, b = explictReturnType(1, 2)
	fmt.Println("explict return type", a, b)

	fmt.Println("var declaration at pkg level", pkgLevelVar1, pkgLevelVar2)
	fmt.Println("var declaration with initial value", varWithInitVal, varWithInitVal2)

	fmt.Println("var declaration can use `:=`, ", k)

	showDefaultValue()

	inferDataTypes()

	useConsts()
}
