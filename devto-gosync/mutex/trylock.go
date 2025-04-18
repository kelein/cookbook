package mutex

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/petermattis/goid"
)

// RecursiveMutex stands for a recursive mutex
type RecursiveMutex struct {
	sync.Mutex
	owner int64
	depth int32
}

// Lock gains the lock
func (m *RecursiveMutex) Lock() {
	gid := goid.Get()
	if atomic.LoadInt64(&m.owner) == gid {
		m.depth++
		return
	}
	m.Mutex.Lock()
	atomic.StoreInt64(&m.owner, gid)
	m.depth = 1
}

// Unlock releases the lock
func (m *RecursiveMutex) Unlock() {
	gid := goid.Get()
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("unlock of wrong owner(%d): %d", m.owner, gid))
	}
	m.depth--
	if m.depth != 0 {
		return
	}
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}

// TokenTryMutex stands for a token try mutex
type TokenTryMutex struct {
	sync.Mutex
	token int64
	depth int32
}

// Lock gains the lock with a token
func (m *TokenTryMutex) Lock(token int64) {
	if atomic.LoadInt64(&m.token) == token {
		m.depth++
		return
	}

	m.Mutex.Lock()
	atomic.StoreInt64(&m.token, token)
	m.depth = 1
}

// Unlock release the lock with a token
func (m *TokenTryMutex) Unlock(token int64) {
	if atomic.LoadInt64(&m.token) != token {
		panic(fmt.Sprintf("unlock of wrong owner(%d): %d", m.token, token))
	}

	m.depth--
	if m.depth != 0 {
		return
	}
	atomic.StoreInt64(&m.token, 0)
	m.Mutex.Unlock()
}
