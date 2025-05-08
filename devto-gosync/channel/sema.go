package channel

import (
	"context"
	"errors"
)

// Semaphore with capacity
type Semaphore struct {
	ch chan struct{}
}

// NewSemaphore create a semaphore with a given capacity
func NewSemaphore(n int) *Semaphore {
	return &Semaphore{
		ch: make(chan struct{}, n),
	}
}

// Acquire blocks until a token is available
func (s *Semaphore) Acquire(ctx context.Context, n int) error {
	if n > cap(s.ch) {
		return errors.New("acquire count exceeds capacity")
	}
	for range n {
		select {
		case s.ch <- struct{}{}:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return nil
}

// Release releases a semaphore token
func (s *Semaphore) Release(ctx context.Context, n int) error {
	if n > cap(s.ch) {
		return errors.New("release count exceeds capacity")
	}
	for range n {
		select {
		case <-s.ch:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return nil
}
