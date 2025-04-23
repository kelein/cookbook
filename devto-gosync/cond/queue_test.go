package cond

import (
	"fmt"
	"log/slog"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestNewLimitQueue(t *testing.T) {
	tests := []struct {
		capacity int
		producer int
		consumer int
	}{
		{10, 5, 5},
		{100, 10, 5},
		{10, 20, 10},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("S%03d/P%03d/C%03d", tt.capacity, tt.producer, tt.consumer)
		t.Run(name, func(t *testing.T) {
			Q := NewLimitQueue(tt.capacity)
			wg := &sync.WaitGroup{}

			for range tt.producer {
				wg.Add(1)
				go func() {
					defer wg.Done()
					time.Sleep(time.Microsecond * time.Duration(rand.Intn(100)))
					value := rand.Intn(1000)
					Q.Push(value)
					slog.Info("pushed to queue", "value", value, "len", Q.Len())
				}()
			}

			for range tt.consumer {
				wg.Add(1)
				go func() {
					defer wg.Done()
					time.Sleep(time.Microsecond * time.Duration(rand.Intn(100)))
					value := Q.Pop()
					slog.Info("popped from queue", "value", value, "len", Q.Len())
				}()
			}

			wg.Wait()
		})
	}
}
