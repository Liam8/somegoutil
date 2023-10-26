package somegoutilities

import "errors"


type MutexQueue struct {
	buffer chan int
}

func NewMutexQueue(capacity int) *MutexQueue {
	return &MutexQueue{
		buffer: make(chan int, capacity),
	}
}

func (r *MutexQueue) Enqueue(item int) error {
	select {
	case r.buffer <- item:
		return nil
	default:
		return errors.New("The queue is full")
	}
}

func (r *MutexQueue) Dequeue() (int, bool) {
	select {
	case item := <- r.buffer:
		return item, true
	default:
		return 0, false
	}
}