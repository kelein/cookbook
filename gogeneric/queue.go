package gogeneric

import "sync"

// GenericQueue queue implemented by the generic
type GenericQueue[T any] struct {
	entry []T
	mu    sync.RWMutex
}

// NewGenericQueue creates a new GenericQueue instance
func NewGenericQueue[T any](capacity int) *GenericQueue[T] {
	return &GenericQueue[T]{entry: make([]T, capacity)}
}

// Size returns the size of the GenericQueue
func (q *GenericQueue[T]) Size() int {
	q.mu.RLock()
	defer q.mu.RUnlock()
	return len(q.entry)
}

// Pop remove front element of the GenericQueue
func (q *GenericQueue[T]) Pop() (_ T) {
	if q.Size() == 0 {
		return
	}
	val := q.Peek()
	q.mu.Lock()
	q.entry = q.entry[1:]
	q.mu.Unlock()
	return val
}

// Peek get the front element of the GenericQueue
func (q *GenericQueue[T]) Peek() (_ T) {
	q.mu.RLock()
	defer q.mu.RUnlock()
	if q.Size() == 0 {
		return
	}
	return q.entry[0]
}

// Push insert element into the GenericQueue
func (q *GenericQueue[T]) Push(v T) {
	q.mu.Lock()
	q.entry = append(q.entry, v)
	q.mu.Unlock()
}
