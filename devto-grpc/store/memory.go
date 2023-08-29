package store

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/copier"

	"cookbook/devto-grpc/repo"
)

// ErrAlreadyExists record already exists
var ErrAlreadyExists = errors.New("record already exists")

// LaptopStore store interface for Laptop
type LaptopStore interface {
	Save(laptop *repo.Laptop) error
	Find(id string) (*repo.Laptop, error)
	Search(ctx context.Context, filter *repo.Filter, got func(laptop *repo.Laptop) error) error
}

// MemoryLaptopStore store Laptop in memory
type MemoryLaptopStore struct {
	data  map[string]*repo.Laptop
	mutex sync.RWMutex
}

// NewMemoryLaptopStore return a MemoryLaptopStore instance
func NewMemoryLaptopStore() *MemoryLaptopStore {
	return &MemoryLaptopStore{
		data: make(map[string]*repo.Laptop),
	}
}

// Find search a laptop by ID
func (m *MemoryLaptopStore) Find(id string) (*repo.Laptop, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	laptop := m.data[id]
	if laptop == nil {
		return nil, nil
	}
	return m.deepClone(laptop)
}

// Save savs the laptop into store
func (m *MemoryLaptopStore) Save(laptop *repo.Laptop) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}
	item, err := m.deepClone(laptop)
	if err != nil {
		return err
	}
	m.data[item.Id] = item
	return nil
}

// Search query laptops with filter conditions
func (m *MemoryLaptopStore) Search(ctx context.Context, filter *repo.Filter, got func(laptop *repo.Laptop) error) error {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for _, laptop := range m.data {
		if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
			log.Print("context canceled while loop store data")
			return nil
		}
		if m.matched(filter, laptop) {
			target, err := m.deepClone(laptop)
			if err != nil {
				log.Printf("deep copy matched laptop failed: %v", err)
				return fmt.Errorf("deep copy error: %w", err)
			}
			if err := got(target); err != nil {
				return fmt.Errorf("search callbak got error: %w", err)
			}
		}
	}
	return nil
}

func (m *MemoryLaptopStore) matched(filter *repo.Filter, laptop *repo.Laptop) bool {
	if laptop.GetPriceUsd() > filter.GetMaxPriceUsed() {
		return false
	}
	if laptop.GetCpu().GetCores() < filter.GetMinCpuCores() {
		return false
	}
	if toBit(laptop.GetRam()) < toBit(filter.GetMinRam()) {
		return false
	}
	return true
}

func (m *MemoryLaptopStore) deepClone(laptop *repo.Laptop) (*repo.Laptop, error) {
	item := &repo.Laptop{}
	err := copier.Copy(item, laptop)
	if err != nil {
		return nil, fmt.Errorf("copy laptop error: %w", err)
	}
	return item, nil
}

// toBit convert memory capcity with uint to bit
func toBit(mem *repo.Memory) uint64 {
	value := mem.GetValue()
	switch mem.GetUnit() {
	case repo.Memory_BIT:
		return value
	case repo.Memory_BYTE:
		// 	8 = 2^3
		return value << 3
	case repo.Memory_KILOBYTE:
		// 8 * 1024 = 2^3 * 2^10
		return value << 13
	case repo.Memory_MEGABYTE:
		return value << 23
	case repo.Memory_GIGABYTE:
		return value << 33
	case repo.Memory_TERABYTE:
		return value << 43
	default:
		return 0
	}
}
