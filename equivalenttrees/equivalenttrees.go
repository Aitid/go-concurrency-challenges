//go:build !solution
// +build !solution

package equivalenttrees

import (
	"golang.org/x/tour/tree"
)

func Walk(root *tree.Tree, ch chan int) {
}

func Same(t1, t2 *tree.Tree) bool {
	return false
}
