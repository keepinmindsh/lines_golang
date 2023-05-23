package tests

import (
	"github.com/reactivex/rxgo/v2"
	"testing"
	"time"
)

func Test_Debounce(t *testing.T) {
	observable := rxgo.Just(1, 2, 3, 4, 5)()

	debounce := observable.Debounce(rxgo.WithDuration(250 * time.Millisecond))

	for item := range debounce.Observe() {
		t.Log(item.V)
	}
}
