package main

import "fmt"

/* Introduction
- golang map is a reference to a hash table
- key must be comparable, int, float64, rune, string,
   comparable array and structure, pointer, etc are OK.
   Noncomparable arrays and structs are not OK.
- you cannot add value to a nil map
*/

type vertex struct {
	x, y int
}

var m map[string]vertex

func mapBasic() {
	m = make(map[string]vertex)
	m["Bob"] = vertex{x: 1, y: 2}
	fmt.Println(m)

	m2 := map[string]vertex{
		"John": vertex{x: 1, y: 2},
		"Wang": vertex{x: 3, y: 4},
	}
	fmt.Println(m2)

	// or shorter:
	m3 := map[string]vertex{
		"John": {x: 1, y: 2},
		"Wang": {x: 3, y: 5},
	}
	fmt.Println(m3)

	name_del := "John"
	delete(m2, name_del)
	elem, ok := m[name_del]
	if !ok {
		fmt.Printf("%s not in m3, elem = %v\n", name_del, elem)
	}

	/* *****NOTE*****
	if you insert into a nil map, a runtime error will be thrown
	var m map[string]int
	m["a"] = 1 // This is not correct
	*/
}

func iterateMap(m map[string]vertex) {
	for idx, val := range m {
		fmt.Printf("m[%s] = [%d]\n", idx, val)
	}
}

func checkKeyExists() {
	m := map[string]float64{"pi": 3.14}

	// method 1: check second return value
	v, found := m["pi"] // v == 3.14  found == true
	v, found = m["pie"] // v == 0.0   found == false
	_, found = m["pi"]  // found == true
	fmt.Println(v, found)

	// method 2: use second return value directly in an if statement
	if v, found := m["pi"]; found {
		fmt.Println(v)
	}
	fmt.Println(v, found)

	// method 3: check for zero value  SHOULDN'T USE THIS
	v = m["pi"]  // v == 3.14
	v = m["pie"] // v == 0.0 (zero value)
	// Warning: This approach doesn't work if the zero value is a possible key.
	// example
	// m := map[int]int,  v := m[0]
	// if return 0, you can't say it's a value 0 for key 0 (key 0 exists)
	// or default zero value for key 0 (key 0 doesn't exist)
}

func main() {
	mapBasic()
	iterateMap(m)
	checkKeyExists()
}
