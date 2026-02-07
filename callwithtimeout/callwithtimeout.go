//go:build !solution
// +build !solution

package callwithtimeout

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Response struct {
	message string
	err     error
}

func ExternalAPIRequest() (string, error) {
	// Simulate a slow API request
	n := rand.Intn(8) + 2
	time.Sleep(time.Duration(n) * time.Second)
	return fmt.Sprintf("data from API. Latency: %d", n), nil
}

func CallWithTimeout(ctx context.Context, fn func() (string, error), timeout time.Duration) (string, error) {
	ch := make(chan Response, 1)
	ctxTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	go func() {
		f, err := fn()
		ch <- Response{message: f, err: err}
	}()
	select {
	case res := <-ch:
		return res.message, res.err
	case <-ctxTimeout.Done():
		return "", ctxTimeout.Err()
	}
}
