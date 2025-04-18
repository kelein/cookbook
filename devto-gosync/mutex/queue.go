package mutex

import "sync"

// SliceQueue stands for a safety slice queue
type SliceQueue struct {
	data []any

	mu   sync.Mutex
	cond *sync.Cond
}

// NewSliceQueue creates a new slice queue instance
func NewSliceQueue(size int) *SliceQueue {
	q := &SliceQueue{data: make([]any, 0, size)}
	q.cond = sync.NewCond(&q.mu)
	return q
}

// Enqueue pushes an element into the queue end
func (q *SliceQueue) Enqueue(item any) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.data = append(q.data, item)
	q.cond.Broadcast()
}

// Dequeue pops an element from the queue head
func (q *SliceQueue) Dequeue() any {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.data) == 0 {
		q.cond.Wait()
		return nil
	}

	item := q.data[0]
	q.data = q.data[1:]
	q.cond.Signal()
	return item
}

// Len return the length of the queue
func (q *SliceQueue) Len() int { return len(q.data) }
