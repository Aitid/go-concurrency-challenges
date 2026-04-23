//go:build !solution
// +build !solution

package callwithtimeout

import (
	"context"
	"time"
)

type Response struct {
	message string
	err     error
}

func ExternalAPIRequest() (string, error) {
	return "", nil
}

func CallWithTimeout(ctx context.Context, fn func() (string, error), timeout time.Duration) (string, error) {
	return "", nil
}
