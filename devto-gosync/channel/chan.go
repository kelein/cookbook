package channel

import (
	"fmt"
	"log/slog"
	"time"
)

func printOrder(count int) {
	type sign struct{}
	chs := make([]chan sign, count)
	for i := range count {
		chs[i] = make(chan sign)
	}

	worker := func(index int, ch, next chan sign) {
		for {
			token := <-ch
			slog.Info("worker", "index", index+1, "char", string('A'+rune(index)))
			time.Sleep(time.Millisecond * 100)
			next <- token
		}
	}

	for i := range count {
		go worker(i, chs[i], chs[(i+1)%count])
	}

	chs[0] <- sign{}
	select {
	case <-time.After(time.Second * 30):
		slog.Info("process timeout and exiting")
	}
}

func printSequence(n int) {
	chs := make([]chan struct{}, n)
	for i := range n {
		chs[i] = make(chan struct{})
	}

	for i := range n {
		go func(i int) {
			for {
				<-chs[i%n]
				time.Sleep(time.Millisecond * 100)
				slog.Info("current goroutine", "i", i, "num", i+1, "letter", string('A'+rune(i)))
				chs[(i+1)%n] <- struct{}{}
			}
		}(i)
	}

	chs[0] <- struct{}{}
	select {
	case <-time.After(time.Second * 30):
		slog.Info("process timeout and exiting")
	}
}

func printNumberLetter() {
	sign := struct{}{}
	done := make(chan struct{})
	number := make(chan struct{})
	letter := make(chan struct{})

	// Number goroutine
	go func() {
		i := 1
		for {
			<-number
			// slog.Info("print", "number", i)
			fmt.Printf("%02d%02d\n", i, i+1)
			// fmt.Printf("%d%d\n", i, i+1)
			time.Sleep(time.Millisecond * 100)
			i += 2
			letter <- sign
		}
	}()

	// Letter goroutine
	go func() {
		i := 'A'
		for {
			select {
			case <-letter:
				if i > 'Z' {
					done <- sign
					return
				}
				// slog.Info("print", "letter", string(i))
				fmt.Printf("%c%c\n", i, i+1)
				time.Sleep(time.Millisecond * 100)
				i += 2
				number <- sign
			}
		}
	}()

	number <- sign
	<-done
}
