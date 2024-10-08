package main

import (
	"log/slog"
	"sync"
	"time"
)

func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
		slog.Info("producer starting", "id", i)
		time.Sleep(time.Millisecond * 200)
	}
	close(ch)
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		slog.Info("consumer processing", "id", i)
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go consumer(ch, &wg)
	go producer(ch, &wg)

	wg.Wait()
}
