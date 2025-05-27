package main

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime/trace"
	"time"
)

func main() {
	f, err := os.Create(filepath.Join(os.TempDir(), "trace.out"))
	if err != nil {
		slog.Error("failed to create trace file", "error", err)
		os.Exit(1)
	}
	slog.Info("trace file created", "path", f.Name())
	defer f.Close()

	if err := trace.Start(f); err != nil {
		slog.Error("failed to start trace", "error", err)
		os.Exit(1)
	}
	defer trace.Stop()

	execute()
}

func execute() {
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
				fmt.Printf("%02d %02d ", i, i+1)
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
				fmt.Printf("%c%c ", i, i+1)
				time.Sleep(time.Millisecond * 100)
				i += 2
				number <- sign
			}
		}
	}()

	number <- sign
	<-quit
}
