package main

import (
	"context"
)

type TokenBucket struct{}

func (t *TokenBucket) Allow() bool {
	return false
}

func (t *TokenBucket) Wait(ctx context.Context) error {
	return nil
}

func main() {
}
