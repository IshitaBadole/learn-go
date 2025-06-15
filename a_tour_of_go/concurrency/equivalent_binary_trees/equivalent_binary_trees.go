package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// A Tour Of Go Exercise: Equivalent Binary Trees
// can be found here: https://go.dev/tour/concurrency/8

// Walk walks the tree t sending all values
// from the tree to the channel ch.
// * function with sender to channel
// * implements in-order tree traversal
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Start goroutines to walk both the trees
	go func() {
		Walk(t1, ch1)
		// close the channel after walking is done
		close(ch1)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	// compare values from both the channels
	for {
		// ok is false if there are no more values to read 
		// or the channel is closed
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		// If values mismatch or one channel finishes early
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		// Both channels done, and all values matched
		if !ok1 && !ok2 {
			return true
		}
	}
}

func main() {
	ch := make(chan int)

	// tree.New(k) returns a binary tree 
	// holding int values k,2k,.. 10k. (so 10 nodes)

	go func(){
		// walk the tree
		Walk(tree.New(1), ch)
		// ch will now hold values 1,2.. 10
		close(ch)
	}()

	// use range to keep reading the channel till it closes
	// recall range-close Go pattern
	for v := range ch {
		fmt.Printf("Value : %d\n", v)
	}

	same := Same(tree.New(1), tree.New(1))
	fmt.Printf("Same trees? %v\n", same)

	same = Same(tree.New(1), tree.New(2))
	fmt.Printf("Same trees? %v\n", same)
}