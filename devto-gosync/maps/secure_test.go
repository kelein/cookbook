package maps

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
)

func TestNewSecureMap(t *testing.T) {
	tests := []struct {
		keynum int
	}{
		{10},
		{100},
		{500},
		{1000},
	}
	for _, tt := range tests {
		t.Run(gofakeit.Zip(), func(t *testing.T) {
			sem := NewSecureMap()
			wg := &sync.WaitGroup{}
			for i := range tt.keynum {
				wg.Add(1)
				go func(i int) {
					defer wg.Done()
					value := fmt.Sprintf("%d-%s", i, strings.Replace(gofakeit.Name(), " ", "-", -1))
					sem.Set(strconv.Itoa(i), value)
				}(i)
			}
			wg.Wait()

			for j := range tt.keynum {
				value, ok := sem.Get(strconv.Itoa(j))
				slog.Info("query key from map", "key", j, "value", value, "ok", ok)
			}
		})
	}
}

func BenchmarkSecureMap(b *testing.B) {
	for range b.N {
		sem := NewSecureMap()
		wg := &sync.WaitGroup{}
		for j := range 1000 {
			wg.Add(1)
			go func(j int) {
				defer wg.Done()
				value := fmt.Sprintf("%d-%s", j, strings.Replace(gofakeit.Name(), " ", "-", -1))
				sem.Set(strconv.Itoa(j), value)
			}(j)
		}
		wg.Wait()

		for k := range 1000 {
			value, ok := sem.Get(strconv.Itoa(k))
			slog.Info("query key from map", "key", k, "value", value, "ok", ok)
		}
	}
}
