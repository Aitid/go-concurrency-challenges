package fanin

import (
	"context"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ch1 := make(chan int)
		ch2 := make(chan int)

		go func() {
			for i := 0; i < 5; i++ {
				ch1 <- i
			}
			close(ch1)
		}()
		go func() {
			for i := 5; i < 10; i++ {
				ch2 <- i
			}
			close(ch2)
		}()
		result := []int{}

		done := make(chan struct{})
		go func() {
			for v := range merge(ctx, ch1, ch2) {
				result = append(result, v)
			}
			close(done)
		}()

		select {
		case <-done:
			assert.ElementsMatch(t, result, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		case <-time.After(2 * time.Second):
			t.Fatal("Happy path timed out! Deadlock in merge.")
		}
	})
	t.Run("zero inputs", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		out := merge[int](ctx)

		select {
		case _, ok := <-out:
			assert.False(t, ok, "channel must be closed immediately")
		case <-time.After(time.Second):
			t.Fatal("merge blocked on zero inputs")
		}
	})
	t.Run("single channel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ch := make(chan int)

		go func() {
			defer close(ch)
			ch <- 42
		}()

		var result []int
		done := make(chan struct{})
		go func() {
			for v := range merge(ctx, ch) {
				result = append(result, v)
			}
			close(done)
		}()

		select {
		case <-done:
			assert.Equal(t, []int{42}, result)
		case <-time.After(2 * time.Second):
			t.Fatal("Single channel timed out! Deadlock in merge.")
		}
	})
	t.Run("slow consumer", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ch := make(chan int)

		go func() {
			defer close(ch)
			for i := 0; i < 100; i++ {
				ch <- i
			}
		}()

		out := merge(ctx, ch)

		time.Sleep(100 * time.Millisecond)

		done := make(chan struct{})
		count := 0
		go func() {
			for range out {
				count++
			}
			close(done)
		}()
		select {
		case <-done:
			assert.Equal(t, 100, count)
		case <-time.After(2 * time.Second):
			t.Fatal("Slow consumer timed out! Deadlock in merge.")
		}
	})
	t.Run("early exit no goroutine leak", func(t *testing.T) {
		before := runtime.NumGoroutine()

		ctx, cancel := context.WithCancel(context.Background())

		ch := make(chan int)

		go func() {
			for {
				ch <- 1
			}
		}()

		out := merge(ctx, ch)

		<-out
		cancel()

		time.Sleep(200 * time.Millisecond)

		after := runtime.NumGoroutine()

		assert.LessOrEqual(
			t,
			after,
			before+2,
			"goroutines leaked",
		)
	})
	t.Run("context cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		ch1 := make(chan int, 10)
		ch1 <- 1

		// We run the check in a loop. Because it's a 50/50 chance, running it 100 times
		// makes the chance of passing (1/2)^100 — virtually zero.
		for i := 0; i < 100; i++ {
			out := merge(ctx, ch1)

			select {
			case v, ok := <-out:
				if ok {
					t.Fatalf("Iteration %d: BUG FOUND! Received %v from a cancelled context", i, v)
				}
			case <-time.After(10 * time.Millisecond):
				t.Fatal("Merge leaked or hung")
			}

			select {
			case ch1 <- 1:
			default:
			}
		}
	})
	t.Run("nil channels", func(t *testing.T) {
		testCtx, testCancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer testCancel()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ch1 := make(chan int, 1)
		ch1 <- 100
		close(ch1)

		out := merge(ctx, nil, ch1, nil)

		select {
		case val, ok := <-out:
			if !ok {
				t.Fatal("expected value 100, but channel closed early")
			}
			assert.Equal(t, 100, val)

			select {
			case _, ok := <-out:
				assert.False(t, ok, "channel should close normally")
			case <-testCtx.Done():
				t.Fatal("Deadlock detected: out channel never closed after processing inputs")
			}

		case <-testCtx.Done():
			t.Fatal("Test timed out: merge function is deadlocked (check WaitGroup logic)")
		}
	})
	t.Run("strict leak test", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		ch1 := make(chan int)
		ch2 := make(chan int)
		ch3 := make(chan int)

		out := merge(ctx, ch1, ch2, ch3)

		time.Sleep(time.Second)

		cancel()

		select {
		case _, ok := <-out:
			if ok {
				t.Fatal("Should not receive data")
			}
		case <-time.After(500 * time.Millisecond):

			t.Fatal("LEAK DETECTED: out channel is still open because workers are stuck")
		}
	})
}
