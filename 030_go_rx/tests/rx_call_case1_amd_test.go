package tests

import (
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_AMDSample(t *testing.T) {
	observable := rxgo.Amb([]rxgo.Observable{
		rxgo.Just(1, 2, 3)(),
		rxgo.Just(4, 5, 6)(),
	})

	items := observable.Observe()

	for item := range items {
		t.Logf("Value : %v", item.V)
	}
}
