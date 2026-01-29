//go:build solution
// +build solution

package queue

import (
	"errors"
	"sync"
)

var (
	ErrQueueFull  = errors.New("queue is full")
	ErrQueueEmpty = errors.New("queue is empty")
)

type Queue interface {
	Push(val int) error // Adds element to back; returns ErrQueueFull if full
	Pop() (int, error)  // Removes and returns front element; returns ErrQueueEmpty if empty
	Peek() (int, error) // Returns front element without removing; returns ErrQueueEmpty if empty
}

type ringQueue struct {
	mu       sync.Mutex
	capacity int
	size     int
	head     int
	tail     int
	data     []int
}

func NewRingQueue(capacity int) (Queue, error) {
	if capacity <= 0 {
		return nil, errors.New("capacity must be positive")
	}
	data := make([]int, capacity)
	return &ringQueue{data: data, capacity: capacity}, nil
}

func (q *ringQueue) Push(val int) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.size >= q.capacity {
		return ErrQueueFull
	}

	q.data[q.tail] = val
	q.tail = (q.tail + 1) % q.capacity
	q.size++
	return nil
}

func (q *ringQueue) Pop() (int, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.size == 0 {
		return -1, ErrQueueEmpty
	}
	v := q.data[q.head]
	q.head = (q.head + 1) % q.capacity
	q.size--
	return v, nil
}

func (q *ringQueue) Peek() (int, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.size == 0 {
		return -1, ErrQueueEmpty
	}
	return q.data[q.head], nil
}
