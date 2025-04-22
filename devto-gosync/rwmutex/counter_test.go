package rwmutex

import (
	"log/slog"
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := &Counter{}
	wg := &sync.WaitGroup{}

	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 100 {
				counter.Incr()
				time.Sleep(time.Millisecond * 50)
			}
		}()
	}
	wg.Wait()

	timer := time.NewTimer(time.Second * 10)
	select {
	case <-timer.C:
		return
	default:
		slog.Info("current counter", "value", counter.Count())
	}
}
