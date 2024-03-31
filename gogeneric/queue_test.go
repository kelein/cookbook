package gogeneric

import (
	"log/slog"
	"sync"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
)

func TestGenericQueue(t *testing.T) {
	tests := []struct {
		name     string
		kind     string
		capacity int
		entries  int
	}{
		{"empty", "int", 0, 100},
		{"no-empty", "string", 0, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.kind {
			case "int":
				q := NewGenericQueue[int](tt.capacity)
				slog.Info("NewGenericQueue", "Info", q, "Size", q.Size())
				q.Pop()
				q.Peek()
				for range tt.entries {
					v := gofakeit.Number(0, 1000)
					q.Push(v)
					slog.Info("GenericQueue", "Push", v, "Peek", q.Peek(), "Size", q.Size())
				}
				for range tt.entries {

				}

				wg := sync.WaitGroup{}
				wg.Add(2)
				go func() {
					defer wg.Done()
					for range tt.entries {
						v := gofakeit.Number(0, 1000)
						q.Push(v)
						slog.Info("GenericQueue", "Push", v, "Peek", q.Peek(), "Size", q.Size())
					}
				}()
				go func() {
					defer wg.Done()
					for range tt.entries {
						slog.Info("GenericQueue", "Pop", q.Pop(), "Size", q.Size())
					}
				}()
				wg.Wait()

			case "string":
				q := NewGenericQueue[string](tt.capacity)
				slog.Info("NewGenericQueue", "Info", q, "Size", q.Size())
				q.Pop()
				q.Peek()

				wg := sync.WaitGroup{}
				wg.Add(2)
				go func() {
					defer wg.Done()
					for range tt.entries {
						v := gofakeit.MiddleName()
						q.Push(v)
						slog.Info("GenericQueue", "Push", v, "Peek", q.Peek(), "Size", q.Size())
					}
				}()
				go func() {
					defer wg.Done()
					for range tt.entries {

						slog.Info("GenericQueue", "Pop", q.Pop(), "Peek", q.Peek(), "Size", q.Size())
					}
				}()
				wg.Wait()
			}
		})
	}
}
