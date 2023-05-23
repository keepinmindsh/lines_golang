package tests

import (
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_ConcatTest(t *testing.T) {
	observable := rxgo.Concat([]rxgo.Observable{
		rxgo.Just(1, 2, 3)(),
		rxgo.Just(4, 5, 6)(),
	})

	for item := range observable.Observe() {
		t.Log(item.V)
	}
}
