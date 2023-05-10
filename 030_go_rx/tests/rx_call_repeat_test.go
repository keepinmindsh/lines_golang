package tests

import (
	"github.com/reactivex/rxgo/v2"
	"testing"
	"time"
)

func Test_Repeat(t *testing.T) {

	observable := rxgo.Just(1, 2, 3)().Repeat(10, rxgo.WithDuration(time.Second))

	items := observable.Observe()

	for item := range items {
		t.Logf("%v", item.V)
	}
}
