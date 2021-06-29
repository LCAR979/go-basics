package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.

func WalkHelper(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	WalkHelper(t.Left, ch)
	ch <- t.Value
	WalkHelper(t.Right, ch)
}

func Walk(t *tree.Tree, ch chan int) {
	WalkHelper(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int, 5)
	c2 := make(chan int, 5)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for i := range c1 {
		if v, hasMore := <-c2; !hasMore || i != v {
			return false
		}
	}
	if _, hasMore := <-c2; hasMore {
		return false
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(2), tree.New(3)))
}
