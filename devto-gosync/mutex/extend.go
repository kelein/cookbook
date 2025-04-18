package mutex

import "sync"

// ExtendMutex extends the sync.Mutex with a trylock method
type ExtendMutex struct {
	sync.Mutex
}

func (m *ExtendMutex) TryLock() bool { return false }
