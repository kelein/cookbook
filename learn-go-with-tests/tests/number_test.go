package tests

import (
	"fmt"
	"sync"
	"testing"
)

func TestNumberLetterPrint(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	numbers := make(chan int)
	letters := make(chan string)

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			numbers <- i
			fmt.Println(<-letters)
		}
	}()

	go func() {
		defer wg.Done()
		for _, letter := range []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"} {
			fmt.Println(<-numbers)
			letters <- letter
		}
	}()

	wg.Wait()
}
