package singleton

import (
	"sync"
	"sync/atomic"
)

var instance = new(Singleton)

// Singleton stands for a singleton object
type Singleton struct{}

// GetInstance return a singleton instance
func GetInstance() *Singleton {
	return instance
}

/* ------------------------------------------------------ */
/*                     Lazy Load Mode                     */
/* ------------------------------------------------------ */

// Object stands for a singleton object
type Object struct{}

var singleObject *Object

// GetObject return a singleton object
func GetObject() *Object {
	if singleObject == nil {
		singleObject = new(Object)
	}
	return singleObject
}

var mtx sync.Mutex

// GetObjectWithLock return a singleton object with mutex
func GetObjectWithLock() *Object {
	mtx.Lock()
	defer mtx.Unlock()

	if singleObject == nil {
		singleObject = new(Object)
	}
	return singleObject
}

/* ------------------------------------------------------ */
/*                       With Atomic                      */
/* ------------------------------------------------------ */

var loaded uint32

// GetObjectWithAtomic return a singleton object with atomic
func GetObjectWithAtomic() *Object {
	if atomic.LoadUint32(&loaded) == 1 {
		return singleObject
	}

	mtx.Lock()
	defer mtx.Unlock()
	if loaded == 0 {
		defer atomic.StoreUint32(&loaded, 1)
		singleObject = new(Object)
	}
	return singleObject
}

/* ------------------------------------------------------ */
/*                     With sync.Once                     */
/* ------------------------------------------------------ */

var once sync.Once

// GetObjectOnce return a singleton with sync.Once
func GetObjectOnce() *Object {
	once.Do(func() { singleObject = new(Object) })
	return singleObject
}

/* ------------------------------------------------------ */
/*                     Once Implement                     */
/* ------------------------------------------------------ */

// MyOnce implement sync Once
type MyOnce struct {
	done uint32
	lock sync.Mutex
}

// Do calls function f only once
func (mo *MyOnce) Do(f func()) {
	if atomic.LoadUint32(&mo.done) == 1 {
		return
	}

	if atomic.LoadUint32(&mo.done) == 0 {
		mo.lock.Lock()
		defer mo.lock.Unlock()
		if mo.done == 0 {
			defer atomic.StoreUint32(&mo.done, 1)
			f()
		}
	}
}

var myOnce MyOnce

// GetObjectWithMyOnce return a singleton with custom Once
func GetObjectWithMyOnce() *Object {
	myOnce.Do(func() { singleObject = new(Object) })
	return singleObject
}
