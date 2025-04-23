package cond

import (
	"log/slog"
	"math/rand"
	"sync"
	"time"
)

func startRunner() {
	ready := 0
	sema := sync.NewCond(&sync.Mutex{})

	for i := range 10 {
		go func(i int) {
			time.Sleep(time.Second * time.Duration(rand.Intn(10)))

			sema.L.Lock()
			ready++
			sema.L.Unlock()

			slog.Info("player in ready state", "index", i)
			sema.Broadcast()
		}(i)
	}

	sema.L.Lock()
	for ready < 10 {
		sema.Wait()
		slog.Info("waiting for players to be ready", "current", ready)
	}
	sema.L.Unlock()

	slog.Info("all players are ready to start running")
}
