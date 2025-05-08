package group

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-kratos/kratos/pkg/sync/errgroup"
)

func TestKratosErrgroup(t *testing.T) {
	var g errgroup.Group
	// 只使用一个goroutine处理子任务
	g.GOMAXPROCS(1)

	var count int64
	g.Go(func(ctx context.Context) error {
		//睡眠5秒，把这个goroutine占住
		time.Sleep(time.Second)
		return nil
	})

	// 并发一万个goroutine执行子任务，
	// 理论上这些子任务都会加入到Group的待处理列表中
	total := 10000
	for range total {
		go func() {
			g.Go(func(ctx context.Context) error {
				atomic.AddInt64(&count, 1)
				return nil
			})
		}()
	}

	// 等待所有的子任务完成。理论上10001个子任务都会被完成
	if err := g.Wait(); err != nil {
		panic(err)
	}

	got := atomic.LoadInt64(&count)
	if got != int64(total) {
		panic(fmt.Sprintf("expect %d but got %d", total, got))
	}
}
