package fanout

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRoundRobinSplit(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		chIn := make(chan int)
		n := 2

		go func() {
			defer close(chIn)
			for i := 0; i < 6; i++ {
				chIn <- i
			}
		}()

		chOut := RoundRobinSplit(chIn, n)

		var wg sync.WaitGroup

		results := make([][]int, n)
		var mu sync.Mutex
		wg.Add(len(chOut))
		go func() {
			for i, ch := range chOut {
				go func() {
					defer wg.Done()
					for v := range ch {
						mu.Lock()
						results[i] = append(results[i], v)
						mu.Unlock()
					}
				}()
			}
		}()
		wg.Wait()

		assert.Equal(t, []int{0, 2, 4}, results[0])
		assert.Equal(t, []int{1, 3, 5}, results[1])
	})
	t.Run("N zero panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("expected panic for n == 0")
			}
		}()
		chIn := make(chan int)

		_ = RoundRobinSplit(chIn, 0)
	})
	t.Run("closes outputs", func(t *testing.T) {
		in := make(chan int)
		outs := RoundRobinSplit(in, 3)

		close(in)

		for i, ch := range outs {
			_, ok := <-ch
			if ok {
				t.Fatalf("output channel %d not closed", i)
			}
		}
	})
	t.Run("unread output blocks", func(t *testing.T) {
		in := make(chan int)
		outs := RoundRobinSplit(in, 2)

		// Only consume from outs[0]
		go func() {
			for range outs[0] {
			}
		}()

		done := make(chan struct{})

		go func() {
			in <- 1
			in <- 2 // routed to outs[1], nobody reads → blocks
			close(done)
		}()

		select {
		case <-done:
			t.Fatalf("expected send to block, but it completed")
		case <-time.After(100 * time.Millisecond):
			// expected
		}
	})
}
