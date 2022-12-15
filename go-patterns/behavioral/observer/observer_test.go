package observer

import (
	"log"
	"testing"
	"time"
)

func TestObserver_Notify(t *testing.T) {
	t.Run("Notify", func(t *testing.T) {
		sub := &Subjector{}
		sub.Subscribe(&MyObserver{})
		sub.Subscribe(&AnotherObserver{})
		sub.Notify("Start Notify")
	})
}

func TestEventBus_Publish(t *testing.T) {
	f1 := func(name, action string) {
		time.Sleep(time.Second)
		log.Printf("%s --> %s", name, action)
	}

	f2 := func(name, action string) {
		time.Sleep(time.Second)
		log.Printf("%s --> %s", name, action)
	}

	t.Run("EventBus", func(t *testing.T) {
		bus := NewAsyncEventBus()
		bus.Subscribe("topic-1", f1)
		bus.Subscribe("topic-2", f2)
		bus.Publish("topic-1", "John", "play")
		bus.Publish("topic-2", "Cathy", "work")
		time.Sleep(time.Second * 3)
	})
}
