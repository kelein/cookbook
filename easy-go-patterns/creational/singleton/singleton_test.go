package singleton

import (
	"log"
	"sync"
	"testing"
)

func TestGetObjectOnce(t *testing.T) {
	obj1 := GetObjectOnce()
	log.Printf("obj: %p", obj1)

	obj2 := GetObjectOnce()
	log.Printf("obj2: %p", obj2)

	if obj1 != obj2 {
		t.Logf("get different objects: %p != %p", obj1, obj2)
	}
}

func TestGetObjectWithAtomic(t *testing.T) {
	t.Run("GetObjectWithAtomic", func(t *testing.T) {
		wg := &sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				got := GetObjectWithAtomic()
				log.Printf("GetObjectWithAtomic() = Object<%p>", got)
			}()
		}
		wg.Wait()
	})
}

func TestGetObjectWithMyOnce(t *testing.T) {
	tests := []struct {
		name string
		num  int
	}{
		{"100-Groutines", 100},
		{"1000-Groutines", 1000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg := &sync.WaitGroup{}
			for i := 0; i < tt.num; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					got := GetObjectWithMyOnce()
					log.Printf("GetObjectWithMyOnce() = Object<%p>", got)
				}()
			}
			wg.Wait()
		})
	}
}

func TestGetObjectWithLock(t *testing.T) {
	tests := []struct {
		name string
		num  int
	}{
		{"100-Groutines", 100},
		{"1000-Groutines", 1000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg := &sync.WaitGroup{}
			for i := 0; i < tt.num; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					got := GetObjectWithLock()
					log.Printf("GetObjectWithLock() = Object<%p>", got)
				}()
			}
			wg.Wait()
		})
	}
}
