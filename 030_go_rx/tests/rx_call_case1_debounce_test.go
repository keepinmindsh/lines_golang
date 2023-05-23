package tests

import (
	"github.com/reactivex/rxgo/v2"
	"sync"
	"testing"
	"time"
)

// FIXME 활용 부분에 대해서 검토 예정
func Test_Debounce(t *testing.T) {
	observable := rxgo.Just(1, 2, 3, 4, 5)()

	var wg sync.WaitGroup

	wg.Add(1)
	debounce := observable.Debounce(rxgo.WithDuration(250 * time.Millisecond))

	for item := range debounce.Observe() {
		t.Log(item.V)
		wg.Done()
	}

	wg.Wait()
}
