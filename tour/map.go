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

func main() {
	mapBasic()
	iterateMap(m)
}
