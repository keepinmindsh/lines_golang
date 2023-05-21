package tests

import (
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_AllSample(t *testing.T) {

	observable := rxgo.Just(1, 2, 3, 4)().
		All(func(i interface{}) bool {
			return i.(int) < 10
		})

	items := observable.
		Observe()

	for item := range items {
		t.Logf("Value : %v", item.V)
	}
}
