package store

import "sync"

// RateStore Latop rate store
type RateStore interface {
	Add(laptopID string, score float64) (*Rate, error)
}

// Rate Laptop rate score
type Rate struct {
	Count uint32
	Sum   float64
}

// MemoryRateStore store laptop rate score in memory
type MemoryRateStore struct {
	lock  sync.Mutex
	rates map[string]*Rate
}

// NewMemoryRateStore create a new MemoryRateStore instance
func NewMemoryRateStore() *MemoryRateStore {
	return &MemoryRateStore{
		rates: make(map[string]*Rate),
	}
}

// Add increments the laptop rate score
func (s *MemoryRateStore) Add(laptopID string, score float64) (*Rate, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	rate := s.rates[laptopID]
	if rate != nil {
		rate.Count++
		rate.Sum += score
	} else {
		rate = &Rate{Count: 1, Sum: score}
	}
	s.rates[laptopID] = rate
	return rate, nil
}
