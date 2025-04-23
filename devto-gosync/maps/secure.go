package maps

import (
	"hash/fnv"
	"sync"
)

// ShardCount stands map size of each shard
const ShardCount = 32

// SecureMap for concurrency
type SecureMap []*SecureShard

// SecureShard shard of SecureMap
type SecureShard struct {
	store map[string]any
	mu    sync.RWMutex
}

// NewSecureMap create a new SecureMap instance
func NewSecureMap() SecureMap {
	m := make(SecureMap, ShardCount)
	for i := range ShardCount {
		m[i] = &SecureShard{store: make(map[string]any)}
	}
	return m
}

// Set set key value to the map
func (s *SecureMap) Set(key string, value any) {
	shard := s.getShard(key)
	shard.mu.Lock()
	shard.store[key] = value
	shard.mu.Unlock()
}

// Get get value by key from the map
func (s *SecureMap) Get(key string) (any, bool) {
	shard := s.getShard(key)
	shard.mu.RLock()
	defer shard.mu.RUnlock()
	value, ok := shard.store[key]
	return value, ok
}

// Delete delete key from the map
func (s *SecureMap) Delete(key string) {
	shard := s.getShard(key)
	shard.mu.Lock()
	defer shard.mu.Unlock()
	delete(shard.store, key)
}

func (s SecureMap) getShard(key string) *SecureShard {
	return s[fnv32(key)%uint32(ShardCount)]
}

func fnv32(key string) uint32 {
	hasher := fnv.New32a()
	hasher.Write([]byte(key))
	return hasher.Sum32()
}
