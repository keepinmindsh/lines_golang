package tests

import (
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_FlatMapSample(t *testing.T) {
	observable := rxgo.Just(1, 2, 3)().FlatMap(func(item rxgo.Item) rxgo.Observable {
		return rxgo.Just(item.V.(int)*10, item.V.(int)*100)()
	})

	for item := range observable.Observe() {
		t.Log(item.V)
	}
}
