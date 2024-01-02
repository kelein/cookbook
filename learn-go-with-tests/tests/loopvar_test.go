package tests

import (
	"fmt"
	"sync"
	"testing"
)

func TestLoopVar(t *testing.T) {
	sl := []int{11, 12, 13, 14, 15}
	var wg sync.WaitGroup
	for i := range sl {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("%d : %d\n", i, sl[i])
			wg.Done()
		}(i)
	}
	wg.Wait()
}
