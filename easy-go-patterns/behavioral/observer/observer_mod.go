package observer

import "log"

// StudentListener of abstract Listener
type StudentListener interface {
	OnTeacherComming()
}

// StudentNotifier of abstract Notifier
type StudentNotifier interface {
	Notify()
	AddListener(l StudentListener)
	RemoveListener(l StudentListener)
}

// ClassMonitor stands for observer
type ClassMonitor struct {
	listeners []StudentListener
}

// AddListener add StudentListener into list
func (c *ClassMonitor) AddListener(listener StudentListener) {
	c.listeners = append(c.listeners, listener)
}

// RemoveListener remove StudentListener from list
func (c *ClassMonitor) RemoveListener(listener StudentListener) {
	for idx, l := range c.listeners {
		if l == listener {
			c.listeners = append(c.listeners[:idx], c.listeners[idx+1:]...)
			break
		}
	}
}

// Notify send notification to all listeners
func (c *ClassMonitor) Notify() {
	for _, l := range c.listeners {
		l.OnTeacherComming()
	}
}

// Student of concrete
type Student struct {
	Name  string
	Event string
}

// NewStudent return a new Student instance
func NewStudent(name, event string) *Student {
	return &Student{Name: name, Event: event}
}

// OnTeacherComming of student
func (s *Student) OnTeacherComming() {
	log.Printf("[%s] stop doing %s", s.Name, s.Event)
}

// Doing of student
func (s *Student) Doing() {
	log.Printf("[%s] is doing %s", s.Name, s.Event)
}
