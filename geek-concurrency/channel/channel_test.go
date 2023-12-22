package channel

import (
	"context"
	"testing"
	"time"
)

func TestChanelPrint(t *testing.T) {
	chs := [4]chan struct{}{
		make(chan struct{}),
		make(chan struct{}),
		make(chan struct{}),
		make(chan struct{}),
	}

	for i := 0; i < 4; i++ {
		go func(i int) {
			for {
				<-chs[i%4]
				t.Logf("current groutine: %d", i)
				time.Sleep(time.Second)
				chs[(i+1)%4] <- struct{}{}
			}
		}(i)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	chs[0] <- struct{}{}
	select {
	case <-ctx.Done():
		t.Logf("context done")
	case <-time.After(time.Second * 10):
		t.Logf("context timeout")
	}
}

func TestChannelPlan(t *testing.T) {
	num := 4
	chs := make([]chan struct{}, num)
	for i := 0; i < num; i++ {
		chs[i] = make(chan struct{})
	}

	for i := 0; i < num; i++ {
		go func(i int) {
			for {
				<-chs[i]
				t.Logf("current groutine: %d", i)
				time.Sleep(time.Second)
				chs[(i+1)%num] <- struct{}{}
			}
		}(i)
	}

	chs[0] <- struct{}{}
	select {
	case <-time.After(time.Second * 10):
		return
	}
}
