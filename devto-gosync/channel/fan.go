package channel

import "sync"

// fanout distributes messages from a single input channel to multiple output channels.
func fanout(in <-chan any, out []chan<- any, async bool) {
	go func() {
		wg := sync.WaitGroup{}
		defer func() {
			wg.Wait()
			for _, ch := range out {
				close(ch)
			}
		}()

		for v := range in {
			for _, ch := range out {
				if async {
					wg.Add(1)
					go func(c chan<- any) {
						c <- v
						wg.Done()
					}(ch)
				} else {
					ch <- v
				}
			}
		}
	}()
}
