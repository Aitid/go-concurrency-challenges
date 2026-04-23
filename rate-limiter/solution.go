//go:build solution
// +build solution

package ratelimiter

import (
	"context"
	"sync"
	"time"
)

type TokenBucket struct {
	mu       sync.Mutex
	capacity float64
	tokens   float64
	rate     float64
	last     time.Time
}

func NewTokenBucket(capacity float64, rate float64) *TokenBucket {
	return &TokenBucket{
		capacity: capacity,
		tokens:   capacity,
		rate:     rate,
		last:     time.Now(),
	}
}

func (t *TokenBucket) refill() {
	now := time.Now()

	t.tokens += now.Sub(t.last).Seconds() * t.rate

	if t.tokens > t.capacity {
		t.tokens = t.capacity
	}
	t.last = now
}

func (t *TokenBucket) Allow() bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.refill()

	if t.tokens >= 1 {
		t.tokens--
		return true
	}
	return false
}

func (t *TokenBucket) Wait(ctx context.Context) error {
	t.mu.Lock()

	for {

		t.refill()

		if t.tokens >= 1 {
			t.tokens--
			t.mu.Unlock()
			return nil
		}

		wait := time.Duration((1 - t.tokens) / t.rate * float64(time.Second))
		t.mu.Unlock()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(wait):
		}

	}
}
