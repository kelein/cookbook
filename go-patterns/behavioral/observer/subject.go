package observer

import "log"

// Subject 发布者
type Subject interface {
	Subscribe(o Observer)
	Cancel(o Observer)
	Notify(msg string)
}

// Observer 观察者
type Observer interface {
	Update(msg string)
}

// Subjector 发布者实现
type Subjector struct {
	observers []Observer
}

// Subscribe of Subjector
func (s *Subjector) Subscribe(o Observer) {
	s.observers = append(s.observers, o)
}

// Cancel of Subjector
func (s *Subjector) Cancel(o Observer) {
	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

// Notify of Subjector
func (s *Subjector) Notify(msg string) {
	for _, o := range s.observers {
		o.Update(msg)
	}
}

// MyObserver 实现Observer
type MyObserver struct{}

// Update of MyObserver
func (o *MyObserver) Update(msg string) {
	log.Printf("MyObserver watched msg changed")
}

// AnotherObserver 实现Observer
type AnotherObserver struct{}

// Update of MyObserver
func (o *AnotherObserver) Update(msg string) {
	log.Printf("AnotherObserver watched msg changed")
}
