package mutex

import (
	"fmt"
	"log/slog"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/petermattis/goid"
)

func gonum() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	sub := strings.TrimPrefix(string(buf[:n]), "goroutine ")
	idField := strings.Fields(sub)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("failed to get goroutine id: %v", err))
	}
	return id
}

func Test_gonum(t *testing.T) {
	wg := sync.WaitGroup{}
	for range 10000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			gid := gonum()
			curr := goid.Get()
			slog.Info("current goroutine", "id", gid, "curr", curr)
		}()
	}
	wg.Wait()
}

func Benchmark_gonum(b *testing.B) {
	wg := sync.WaitGroup{}
	for range b.N {
		wg.Add(1)
		go func() {
			defer wg.Done()
			gid := gonum()
			curr := goid.Get()
			slog.Info("current goroutine", "id", gid, "curr", curr)
		}()
	}
	wg.Wait()
	b.ReportAllocs()
}

func TestRecursiveMutex(t *testing.T) {
	count := 0
	wg := sync.WaitGroup{}
	mu := &RecursiveMutex{}
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
