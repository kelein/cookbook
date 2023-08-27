package store

import (
	"errors"
	"fmt"
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

func (m *MemoryLaptopStore) deepClone(laptop *repo.Laptop) (*repo.Laptop, error) {
	item := &repo.Laptop{}
	err := copier.Copy(item, laptop)
	if err != nil {
		return nil, fmt.Errorf("copy laptop error: %w", err)
	}
	return item, nil
}
