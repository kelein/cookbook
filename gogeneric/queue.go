package gogeneric

// GenericQueue queue implemented by the generic
type GenericQueue[T any] []T

// func NewGenericQueue() *GenericQueue {
// 	queue := make([]any, 0)
// 	return queue
// }

// Size returns the size of the GenericQueue
func (q *GenericQueue[T]) Size() int {
	return len(*q)
}

// Pop remove front element of the GenericQueue
func (q *GenericQueue[T]) Pop() (_ T) {
	if q.Size() == 0 {
		return
	}
	val := q.Peek()
	*q = (*q)[1:]
	return val
}

// Peek get the front element of the GenericQueue
func (q *GenericQueue[T]) Peek() (_ T) {
	if q.Size() == 0 {
		return
	}
	return (*q)[0]
}

// Push insert element into the GenericQueue
func (q *GenericQueue[T]) Push(v T) {
	*q = append(*q, v)
}
