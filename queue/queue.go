//go:build !solution
// +build !solution

package queue

import (
	"errors"
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

type ringQueue struct{}

func NewRingQueue(capacity int) (Queue, error) {
	return &ringQueue{}, nil
}

func (q *ringQueue) Push(val int) error {
	return nil
}

func (q *ringQueue) Pop() (int, error) {
	return -1, nil
}

func (q *ringQueue) Peek() (int, error) {
	return -1, nil
}
