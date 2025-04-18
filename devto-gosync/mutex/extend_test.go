package mutex

import (
	"log/slog"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/petermattis/goid"
)

func TestExtendMutex_TryLock(t *testing.T) {
	mu := &ExtendMutex{}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		mu.Lock()
		defer func() {
			mu.Unlock()
			wg.Done()
		}()

		time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	}()

	time.Sleep(time.Second)
	if mu.TryLock() {
		slog.Info("TryLock succeeded")
		mu.Unlock()
		return
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for mu.TryLock() {
			slog.Info("TryLock in spin succeeded")
			mu.Unlock()
			return
		}
	}()

	wg.Wait()
	slog.Info("TryLock failed")
}

func TestExtendMutex_Count(t *testing.T) {
	mu := &ExtendMutex{}
	wg := sync.WaitGroup{}

	for i := range 1000 {
		wg.Add(1)
		go func(i int) {
			defer func() {
				slog.Info("task done", "idx", i, "goid", goid.Get())
				wg.Done()
			}()
			mu.Lock()
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(50)))
			mu.Unlock()
		}(i)
	}

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for range ticker.C {
			slog.Info("current mutex state", "waiting", mu.Count(), "locked", mu.IsLocked(), "woken", mu.IsWoken(), "starving", mu.IsStarving())
		}
	}()

	wg.Wait()
}
