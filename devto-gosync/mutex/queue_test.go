package mutex

import (
	"fmt"
	"log/slog"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestSliceQueue_Enqueue_Dequeue(t *testing.T) {
	tests := []struct {
		size      int
		producers int
		consumers int
	}{
		{10, 1, 2},
	}

	for _, tc := range tests {
		name := fmt.Sprintf("S%d/P%02d/C%02d", tc.size, tc.producers, tc.consumers)
		t.Run(name, func(t *testing.T) {
			Q := NewSliceQueue(tc.size)
			wg := &sync.WaitGroup{}

			for i := range tc.producers {
				wg.Add(1)
				go func(i int) {
					defer wg.Done()

					for Q.Len() < tc.size {
						v := i + rand.Intn(100)
						Q.Enqueue(v)
						slog.Info("produced", "index", i, "value", v, "len", Q.Len())
					}
				}(i)
			}

			for j := range tc.consumers {
				wg.Add(1)
				go func(j int) {
					defer wg.Done()

					v := Q.Dequeue()
					time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
					slog.Info("consumed", "index", j, "value", v, "len", Q.Len())

				}(j)
			}

			wg.Wait()
		})
	}
}
