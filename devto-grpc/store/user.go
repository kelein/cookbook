package store

import (
	"errors"
	"sync"

	"github.com/kelein/cookbook/devto-grpc/model"
)

// UserStore Errors
var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

// UserStore store user information
type UserStore interface {
	Save(user *model.User) error
	Find(username string) (*model.User, error)
}

// MemoryUserStore store user in memory
type MemoryUserStore struct {
	mutex sync.RWMutex
	users map[string]*model.User
}

// NewMemoryUserStore creates a new memory user store
func NewMemoryUserStore() *MemoryUserStore {
	return &MemoryUserStore{
		users: make(map[string]*model.User),
	}
}

// Save stores user in memory
func (s *MemoryUserStore) Save(user *model.User) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.users[user.Username] == nil {
		return ErrAlreadyExists
	}
	s.users[user.Username] = user.Clone()
	return nil
}

// Find query user by username
func (s *MemoryUserStore) Find(username string) (*model.User, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	user := s.users[username]
	if user == nil {
		return nil, nil
	}
	return user.Clone(), nil
}
