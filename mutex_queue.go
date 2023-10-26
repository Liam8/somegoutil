package somegoutil

import "errors"


type MutexQueue[T any] struct {
	buffer chan T
}

func NewMutexQueue[T any](capacity int) *MutexQueue[T] {
	return &MutexQueue[T]{
		buffer: make(chan T, capacity),
	}
}

func (r *MutexQueue[T]) Enqueue(item T) error {
	select {
	case r.buffer <- item:
		return nil
	default:
		return errors.New("The queue is full")
	}
}

func (r *MutexQueue[T]) Dequeue() (T, bool) {
	select {
	case item := <- r.buffer:
		return item, true
	default:
		return *new(T), false
	}
}