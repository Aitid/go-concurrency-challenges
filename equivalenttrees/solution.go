//go:build solution
// +build solution

package equivalenttrees

import (
	"golang.org/x/tour/tree"
)

func Walk(root *tree.Tree, ch chan int) {
	stack := []*tree.Tree{}
	node := root

	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		ch <- node.Value

		node = node.Right
	}
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		val1, ok1 := <-ch1
		val2, ok2 := <-ch2
		if ok1 && ok2 && val1 != val2 {
			return false
		}

		if ok1 && !ok2 || !ok1 && ok2 {
			return false
		}
		if !ok1 && !ok2 {
			return true
		}
	}
}
