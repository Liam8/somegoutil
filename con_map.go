package somegoutil

import (
	"sync"
)

type ConMap[K comparable, V any] struct {
	data map[K]V
	mu sync.RWMutex
}

func NewConMap[K comparable, V any]() *ConMap[K, V] {
	return &ConMap[K, V] {
		data: make(map[K]V),
	}
}

func (r *ConMap[K, V]) Put(key K, val V) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[key] = val
}

func (r *ConMap[K, V]) Get(key K) (V, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	val, ok := r.data[key]
	return val, ok
}