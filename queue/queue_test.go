package queue

import (
	"sync"
	"testing"
)

func TestPushPop(t *testing.T) {
	t.Run("push until full", func(t *testing.T) {
		q, _ := NewRingQueue(2)

		if err := q.Push(1); err != nil {
			t.Fatal(err)
		}
		if err := q.Push(2); err != nil {
			t.Fatal(err)
		}
		if err := q.Push(3); err == nil {
			t.Fatal("expected error when queue is full")
		}
	})

	t.Run("pop until empty", func(t *testing.T) {
		q, _ := NewRingQueue(2)

		q.Push(1)
		q.Push(2)

		v, _ := q.Pop()
		if v != 1 {
			t.Fatalf("expected 1, got %d", v)
		}

		v, _ = q.Pop()
		if v != 2 {
			t.Fatalf("expected 2, got %d", v)
		}

		if _, err := q.Pop(); err == nil {
			t.Fatal("expected error when queue is empty")
		}
	})
}

func TestPeek(t *testing.T) {
	t.Run("peek empty queue", func(t *testing.T) {
		q, _ := NewRingQueue(1)

		if _, err := q.Peek(); err == nil {
			t.Fatal("expected error on empty queue")
		}
	})

	t.Run("peek does not remove item", func(t *testing.T) {
		q, _ := NewRingQueue(1)
		q.Push(42)

		v, _ := q.Peek()
		if v != 42 {
			t.Fatalf("expected 42, got %d", v)
		}

		v, _ = q.Pop()
		if v != 42 {
			t.Fatalf("expected 42 after peek, got %d", v)
		}
	})
}

func TestConcurrentPushPop_Race(t *testing.T) {
	q, _ := NewRingQueue(10)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := range 100 {
			q.Push(i)
		}
	}()
	go func() {
		defer wg.Done()
		for range 100 {
			_, _ = q.Pop()
		}
	}()

	wg.Wait()
}
