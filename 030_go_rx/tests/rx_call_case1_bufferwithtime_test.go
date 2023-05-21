package tests

import (
	"github.com/reactivex/rxgo/v2"
	"sync"
	"testing"
	"time"
)

func Test_BufferWithTime(t *testing.T) {
	ch := make(chan rxgo.Item, 1)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		i := 0
		for range time.Tick(time.Millisecond) {
			ch <- rxgo.Of(i)
			i++

			if i > 5000 {
				close(ch)
				wg.Done()
				break
			}
		}
	}()

	select {
	case <-ch:
		observable := rxgo.FromChannel(ch).BufferWithTime(rxgo.WithDuration(2 * time.Second))

		for item := range observable.Observe() {
			t.Logf("Value : %v", item)
		} // 2.004948752s
	}

	wg.Wait()
}
