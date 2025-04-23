package once

import (
	"sync"
	"sync/atomic"
)

// ErrorOnce extend once with error check
type ErrorOnce struct {
	mu   sync.Mutex
	done uint32
}

// Do extend once.Do with error check
func (o *ErrorOnce) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 1 {
		return nil
	}
	return o.slowDo(f)
}

func (o *ErrorOnce) slowDo(f func() error) error {
	o.mu.Lock()
	defer o.mu.Unlock()

	var err error
	if o.done == 0 {
		err = f()
		if err == nil {
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}

// Done check if the once has been executed
func (o *ErrorOnce) Done() bool {
	o.mu.Lock()
	defer o.mu.Unlock()
	return atomic.LoadUint32(&o.done) == 1
}
