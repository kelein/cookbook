package gogeneric

// GenericQueue queue implemented by the generic
type GenericQueue[T any] struct {
	entry []T
}

// NewGenericQueue creates a new GenericQueue instance
func NewGenericQueue[T any](capacity int) *GenericQueue[T] {
	return &GenericQueue[T]{entry: make([]T, capacity)}
}

// Size returns the size of the GenericQueue
func (q *GenericQueue[T]) Size() int {
	return len(q.entry)
}

// Pop remove front element of the GenericQueue
func (q *GenericQueue[T]) Pop() (_ T) {
	if q.Size() == 0 {
		return
	}
	val := q.Peek()
	q.entry = q.entry[1:]
	return val
}

// Peek get the front element of the GenericQueue
func (q *GenericQueue[T]) Peek() (_ T) {
	if q.Size() == 0 {
		return
	}
	return q.entry[0]
}

// Push insert element into the GenericQueue
func (q *GenericQueue[T]) Push(v T) {
	q.entry = append(q.entry, v)
}
