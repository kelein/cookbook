package optional

import "sync"

// Cache of Local
type Cache struct {
	bucketCount uint64
	locks       []sync.RWMutex
	close       chan struct{}
}

// Option of cache settings
type Option func(*Cache)

// NewCache create a Cache with given options
func NewCache(opts ...Option) *Cache {
	c := &Cache{close: make(chan struct{})}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// WithBucketCount set cache bucket count
func WithBucketCount(n uint64) Option {
	return func(c *Cache) {
		c.bucketCount = n
	}
}
