package main

import (
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	for i := range 90000 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for range 100 {
				go printNumberLetter()
			}
			time.Sleep(time.Second * time.Duration(rand.Intn(i%7+1)))
		}(i)
	}

	wg.Wait()
}

func printNumberLetter() {
	sign := struct{}{}
	quit := make(chan struct{})
	number := make(chan struct{})
	letter := make(chan struct{})

	// * Number Print
	go func() {
		i := 1
		for {
			select {
			case <-number:
				// fmt.Printf("%02d %02d ", i, i+1)
				// slog.Info("printing numbers", "first", i, "second", i+1)
				time.Sleep(time.Millisecond * 100)
				i += 2
				letter <- sign
			}
		}
	}()

	// * Letter Print
	go func() {
		i := 'A'
		for {
			select {
			case <-letter:
				if i > 'Z' {
					quit <- sign
					return
				}
				// fmt.Printf("%c%c ", i, i+1)
				// slog.Info("printing letters", "first", fmt.Sprintf("%c", i), "second", fmt.Sprintf("%c", i+1))
				time.Sleep(time.Millisecond * 100)
				i += 2
				number <- sign
			}
		}
	}()

	number <- sign
	<-quit
}
