//go:build solution
// +build solution

package callwithtimeout

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func ExternalAPIRequest() (string, error) {
	// Simulate a slow API request
	n := rand.Intn(8) + 2
	time.Sleep(time.Duration(n) * time.Second)
	return fmt.Sprintf("data from API. Latency: %d", n), nil
}

type Fn func() (string, error)

func CallWithTimeout(ctx context.Context, fn Fn, timeout time.Duration) (string, error) {
}
