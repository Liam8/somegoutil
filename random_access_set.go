package somegoutil

import "math/rand"

type RandSet[T comparable] struct {
	arr    []T
	idxMap map[T]int
}

func NewRandSet[T comparable]() *RandSet[T] {
	return &RandSet[T]{
		arr:    []T{},
		idxMap: map[T]int{},
	}
}

func (r *RandSet[T]) Add(val T) {
	_, ok := r.idxMap[val]
	if ok {
		return
	}

	r.arr = append(r.arr, val)
	r.idxMap[val] = len(r.arr) - 1
}

func (r *RandSet[T]) Remove(val T) {
	idx, ok := r.idxMap[val]
	if !ok {
		return
	}

	lastIdx := len(r.arr) - 1
	r.arr[idx] = r.arr[lastIdx]
	r.idxMap[r.arr[idx]] = idx

	r.arr = r.arr[:len(r.arr)-1]
	delete(r.idxMap, val)
}

func (r *RandSet[T]) RandomPick() T {
	if len(r.arr) > 0 {
		return r.arr[rand.Intn(len(r.arr))]
	} else {
		return *new(T)
	}
}

func (r *RandSet[T]) Exists(val T) bool {
	_, ok := r.idxMap[val]
	return ok
}
