package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func WalkBranch(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		WalkBranch(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		WalkBranch(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkBranch(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	var solution1 string
	var solution2 string

	for i := range ch1 {
		solution1 += fmt.Sprintf("%v", i)
	}

	for i := range ch2 {
		solution2 += fmt.Sprintf("%v", i)
	}

	return solution1 == solution2
}

func main() {
	res := Same(tree.New(1), tree.New(1))
	fmt.Println(res)
	res = Same(tree.New(1), tree.New(2))
	fmt.Println(res)
}
