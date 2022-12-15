package observer

import (
	"fmt"
	"log"
	"reflect"
	"sync"
)

// EventBus of abstract
type EventBus interface {
	Subscribe(topic string, handler any) error
	Publish(topic string, args ...any)
}

// AsyncEventBus implement EventBus interface
type AsyncEventBus struct {
	handlers map[string][]reflect.Value
	lock     sync.RWMutex
}

// NewAsyncEventBus create a AsyncEventBus instance
func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		handlers: make(map[string][]reflect.Value),
		lock:     sync.RWMutex{},
	}
}

// Subscribe of AsyncEventBus
func (ab *AsyncEventBus) Subscribe(topic string, handler any) error {
	ab.lock.Lock()
	defer ab.lock.Unlock()

	v := reflect.ValueOf(handler)
	if v.Type().Kind() != reflect.Func {
		return fmt.Errorf("invalid handlers function")
	}

	handlers, ok := ab.handlers[topic]
	if !ok {
		handlers = []reflect.Value{}
	}
	ab.handlers[topic] = append(handlers, v)
	return nil
}

// Publish of AsyncEventBus
func (ab *AsyncEventBus) Publish(topic string, args ...any) {
	handlers, ok := ab.handlers[topic]
	if !ok {
		log.Printf("not found handler for topic: %q", topic)
		return
	}

	params := make([]reflect.Value, len(args))
	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}

	for _, f := range handlers {
		go f.Call(params)
	}
}
