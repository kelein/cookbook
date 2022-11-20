package observer

import (
	"log"
	"strings"
	"testing"
)

func TestClassMonitor_Notify(t *testing.T) {
	t.Run("Notify", func(t *testing.T) {
		s1 := NewStudent("John", "Writing")
		s2 := NewStudent("Peter", "Playing")
		s3 := NewStudent("Frank", "Watching")

		log.Print(strings.Repeat("-", 25))
		s1.Doing()
		s2.Doing()
		s3.Doing()

		log.Print(strings.Repeat("-", 25))
		cm := new(ClassMonitor)
		cm.AddListener(s1)
		cm.AddListener(s2)
		cm.AddListener(s3)

		cm.Notify()

		cm.RemoveListener(s3)
	})

}
