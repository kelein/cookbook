package mutex

import (
	"log/slog"
	"sync"
	"testing"
)

func mutexAdd() {
	count := 0
	mu := &sync.Mutex{}
	wg := sync.WaitGroup{}

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	slog.Info("count", "value", count)
}

func TestMutexAdd(t *testing.T) {
	mutexAdd()
}

func BenchmarkMutexAdd(b *testing.B) {
	for range b.N {
		mutexAdd()
	}
	b.ReportAllocs()
}

type Counter struct {
	sync.Mutex
	Count int64
}

func TestCounter(t *testing.T) {
	counter := &Counter{}
	wg := sync.WaitGroup{}

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.Lock()
				counter.Count++
				counter.Unlock()
			}
		}()
	}
	wg.Wait()
	slog.Info("counter", "value", counter.Count)
}

func BenchmarkCounter(b *testing.B) {
	for range b.N {
		counter := &Counter{}
		wg := sync.WaitGroup{}

		for range 10 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for range 1000 {
					counter.Lock()
					counter.Count++
					counter.Unlock()
				}
			}()
		}
		wg.Wait()
		slog.Info("counter", "value", counter.Count)
	}
	b.ReportAllocs()
}

type AdvancedCounter struct {
	mu    sync.Mutex
	count int64
}

func (ac *AdvancedCounter) Incr() {
	ac.mu.Lock()
	ac.count++
	ac.mu.Unlock()
}

func (ac *AdvancedCounter) Count() int64 {
	ac.mu.Lock()
	defer ac.mu.Unlock()
	return ac.count
}

func TestAdvancedCounter(t *testing.T) {
	wg := sync.WaitGroup{}
	counter := &AdvancedCounter{}
	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.Incr()
			}
		}()
	}
	wg.Wait()
	slog.Info("AdvancedCounter", "value", counter.Count())
}

func BenchmarkAdvancedCounter(b *testing.B) {
	for range b.N {
		wg := sync.WaitGroup{}
		counter := &AdvancedCounter{}
		for range 10 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for range 1000 {
					counter.Incr()
				}
			}()
		}
		wg.Wait()
		slog.Info("AdvancedCounter", "value", counter.Count())
	}
	b.ReportAllocs()
	b.ReportMetric(float64(b.N), "ops/s")
}
