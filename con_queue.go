package somegoutil

import (
	"errors"
	"time"
)

type ConQueue[T any] struct {
	buffer chan T
}

func NewConQueue[T any](capacity int) *ConQueue[T] {
	return &ConQueue[T]{
		buffer: make(chan T, capacity),
	}
}

func (r *ConQueue[T]) Enqueue(item T) error {
	select {
	case r.buffer <- item:
		return nil
	default:
		return errors.New("The queue is full")
	}
}

func (r *ConQueue[T]) Dequeue() (T, bool) {
	select {
	case item := <-r.buffer:
		return item, true
	default:
		return *new(T), false
	}
}

func (r *ConQueue[T]) DequeueWithBlock(timeout int) (T, bool) {
	select {
	case item := <-r.buffer:
		return item, true
	case <-time.After(time.Duration(timeout) * time.Millisecond):
		return *new(T), false
	}
}
