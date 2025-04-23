package cond

import (
	"container/list"
	"sync"
)

// LimitQueue stands for a queue with limit size
type LimitQueue struct {
	store *list.List
	cond  *sync.Cond
	size  int
}

// NewLimitQueue create a new limit queue instance
func NewLimitQueue(capacity int) *LimitQueue {
	return &LimitQueue{
		size:  capacity,
		store: list.New(),
		cond:  sync.NewCond(&sync.Mutex{}),
	}
}

// Push adds an element to the queue
func (q *LimitQueue) Push(v any) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	for q.store.Len() >= q.size {
		q.cond.Wait()
	}

	q.store.PushBack(v)
	q.cond.Broadcast()
}

// Pop removes an element from the queue
func (q *LimitQueue) Pop() any {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	for q.store.Len() == 0 {
		q.cond.Wait()
	}

	e := q.store.Front()
	q.store.Remove(e)
	q.cond.Broadcast()
	return e.Value
}

// Len return the length of current queue
func (q *LimitQueue) Len() int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	return q.store.Len()
}
