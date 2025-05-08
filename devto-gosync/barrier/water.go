package barrier

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"sort"
	"sync"
	"time"

	"github.com/marusama/cyclicbarrier"
	"golang.org/x/sync/semaphore"
)

// H2O stands for water molecule
type H2O struct {
	semaO   *semaphore.Weighted
	semaH   *semaphore.Weighted
	barrier cyclicbarrier.CyclicBarrier
}

// NewH2O create a new H2O instance
func NewH2O() *H2O {
	return &H2O{
		semaO:   semaphore.NewWeighted(1),
		semaH:   semaphore.NewWeighted(2),
		barrier: cyclicbarrier.New(3),
	}
}

func (h2o *H2O) hydrogen(releaseHydrogen func()) {
	h2o.semaH.Acquire(context.Background(), 1)

	releaseHydrogen()

	h2o.barrier.Await(context.Background())
	h2o.semaH.Release(1)
}

func (h2o *H2O) oxygen(releaseOxygen func()) {
	h2o.semaO.Acquire(context.Background(), 1)

	releaseOxygen()

	h2o.barrier.Await(context.Background())
	h2o.semaO.Release(1)
}

// WaterFactory produces water molecules of specified count
func WaterFactory(count int) error {
	if count <= 0 {
		err := errors.New("molecule count must be greater than 0")
		slog.Error(err.Error())
		return err
	}

	h2o := NewH2O()
	wg := &sync.WaitGroup{}
	ch := make(chan string, count*3)
	releaseOxygen := func() {
		ch <- "O"
		slog.Info("O")
	}
	releaseHydrogen := func() {
		ch <- "H"
		slog.Info("H")
	}

	// * Hydrogen processes
	for range count * 2 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * time.Duration(rand.IntN(100)))
			h2o.hydrogen(releaseHydrogen)
			slog.Info("hydrogen process step done")
		}()
	}

	// * Oxygen processes
	for range count {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * time.Duration(rand.IntN(100)))
			h2o.oxygen(releaseOxygen)
			slog.Info("oxygen process step done")
		}()
	}

	wg.Wait()

	if len(ch) != count*3 {
		err := errors.New("atoms count mismatch")
		slog.Error(err.Error(), "expected", count*3, "actual", len(ch))
		return err
	}

	s := make([]string, 3)
	for i := range count {
		s[0] = <-ch
		s[1] = <-ch
		s[2] = <-ch
		sort.Strings(s)

		water := s[0] + s[1] + s[2]
		if water != "HHO" {
			err := errors.New("water molecule mismatch")
			slog.Error(err.Error(), "expected", "HHO", "actual", water)
			return err
		}
		slog.Info("water compound done", "molecule", water, "id", fmt.Sprintf("%06d", i))
	}
	return nil
}
