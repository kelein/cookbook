package rwmutex

import "sync"

// Counter for concurrent data count
type Counter struct {
	count uint64
	mu    sync.RWMutex
}

// Incr increments the counter
func (c *Counter) Incr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// Count return the current count
func (c *Counter) Count() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}
