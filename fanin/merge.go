package fanin

import (
	"context"
)

// Merge is a stub. Implement this function to pass all tests.
func merge[T any](ctx context.Context, chs ...<-chan T) <-chan T {
	out := make(chan T)
	close(out)
	return out
}
