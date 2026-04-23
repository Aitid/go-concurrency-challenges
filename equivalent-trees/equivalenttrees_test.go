package equivalenttrees

import (
	"reflect"
	"testing"
	"time"

	"golang.org/x/tour/tree"
)

func collect(ch chan int) []int {
	var result []int
	for v := range ch {
		result = append(result, v)
	}
	return result
}

func TestWalk(t *testing.T) {
	run := func(t *testing.T, root *tree.Tree, expected []int) {
		ch := make(chan int)
		go Walk(root, ch)

		done := make(chan struct{})
		var result []int

		go func() {
			result = collect(ch)
			close(done)
		}()

		select {
		case <-done:
			if !reflect.DeepEqual(result, expected) {
				t.Fatalf("expected %v, got %v", expected, result)
			}
		case <-time.After(2 * time.Second):
			t.Fatal("deadlock: Walk did not close channel")
		}
	}

	t.Run("Empty Tree", func(t *testing.T) {
		run(t, nil, nil)
	})

	t.Run("Single Node", func(t *testing.T) {
		root := &tree.Tree{Value: 42}
		run(t, root, []int{42})
	})

	t.Run("Left Skewed Tree", func(t *testing.T) {
		root := &tree.Tree{
			Value: 3,
			Left: &tree.Tree{
				Value: 2,
				Left:  &tree.Tree{Value: 1},
			},
		}
		run(t, root, []int{1, 2, 3})
	})

	t.Run("Right Skewed Tree", func(t *testing.T) {
		root := &tree.Tree{
			Value: 1,
			Right: &tree.Tree{
				Value: 2,
				Right: &tree.Tree{Value: 3},
			},
		}
		run(t, root, []int{1, 2, 3})
	})

	t.Run("Balanced Tree", func(t *testing.T) {
		root := &tree.Tree{
			Value: 4,
			Left: &tree.Tree{
				Value: 2,
				Left:  &tree.Tree{Value: 1},
				Right: &tree.Tree{Value: 3},
			},
			Right: &tree.Tree{
				Value: 6,
				Left:  &tree.Tree{Value: 5},
				Right: &tree.Tree{Value: 7},
			},
		}
		run(t, root, []int{1, 2, 3, 4, 5, 6, 7})
	})

	t.Run("Duplicate Values", func(t *testing.T) {
		root := &tree.Tree{
			Value: 2,
			Left:  &tree.Tree{Value: 1},
			Right: &tree.Tree{Value: 2},
		}
		run(t, root, []int{1, 2, 2})
	})
}

func TestSame(t *testing.T) {
	t.Run("Both Nil", func(t *testing.T) {
		if !Same(nil, nil) {
			t.Fatal("expected true for two nil trees")
		}
	})

	t.Run("One Nil One NonNil", func(t *testing.T) {
		t1 := &tree.Tree{Value: 1}

		if Same(t1, nil) {
			t.Fatal("expected false when one tree is nil")
		}

		if Same(nil, t1) {
			t.Fatal("expected false when one tree is nil")
		}
	})

	t.Run("Same Single Node", func(t *testing.T) {
		t1 := &tree.Tree{Value: 42}
		t2 := &tree.Tree{Value: 42}

		if !Same(t1, t2) {
			t.Fatal("expected true for identical single-node trees")
		}
	})

	t.Run("Different Single Node Values", func(t *testing.T) {
		t1 := &tree.Tree{Value: 1}
		t2 := &tree.Tree{Value: 2}

		if Same(t1, t2) {
			t.Fatal("expected false for different single-node values")
		}
	})

	t.Run("Same Values Different Structure", func(t *testing.T) {
		// t1 balanced
		t1 := &tree.Tree{
			Value: 2,
			Left:  &tree.Tree{Value: 1},
			Right: &tree.Tree{Value: 3},
		}

		// t2 right-skewed but same inorder
		t2 := &tree.Tree{
			Value: 1,
			Right: &tree.Tree{
				Value: 2,
				Right: &tree.Tree{
					Value: 3,
				},
			},
		}

		if !Same(t1, t2) {
			t.Fatal("expected true for trees with same inorder values")
		}
	})

	t.Run("Different Sizes", func(t *testing.T) {
		t1 := &tree.Tree{
			Value: 2,
			Left:  &tree.Tree{Value: 1},
		}

		t2 := &tree.Tree{
			Value: 2,
			Left:  &tree.Tree{Value: 1},
			Right: &tree.Tree{Value: 3},
		}

		if Same(t1, t2) {
			t.Fatal("expected false for trees of different sizes")
		}
	})

	t.Run("Different Values Same Size", func(t *testing.T) {
		t1 := &tree.Tree{
			Value: 2,
			Left:  &tree.Tree{Value: 1},
			Right: &tree.Tree{Value: 3},
		}

		t2 := &tree.Tree{
			Value: 2,
			Left:  &tree.Tree{Value: 1},
			Right: &tree.Tree{Value: 4},
		}

		if Same(t1, t2) {
			t.Fatal("expected false for trees with different values")
		}
	})

	t.Run("Duplicate Values", func(t *testing.T) {
		t1 := &tree.Tree{
			Value: 2,
			Left:  &tree.Tree{Value: 2},
			Right: &tree.Tree{Value: 3},
		}

		t2 := &tree.Tree{
			Value: 2,
			Left:  &tree.Tree{Value: 2},
			Right: &tree.Tree{Value: 3},
		}

		if !Same(t1, t2) {
			t.Fatal("expected true for trees with duplicate values")
		}
	})
}
