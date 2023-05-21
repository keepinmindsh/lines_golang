package tests

import (
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_AverageSample(t *testing.T) {
	observable := rxgo.Just(1, 2, 3, 4, 6, 7, 8, 9, 10)().AverageInt()

	items := observable.Observe()

	for item := range items {
		t.Logf("Value : %v", item.V)
	}
}
