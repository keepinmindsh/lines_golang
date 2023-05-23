package tests

import (
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_CombineLatest(t *testing.T) {

	observable := rxgo.CombineLatest(func(i ...interface{}) interface{} {
		sum := 0
		for _, v := range i {
			if v == nil {
				continue
			}
			sum += v.(int)
		}
		return sum
	}, []rxgo.Observable{
		rxgo.Just(1, 2)(),
		rxgo.Just(10, 11)(),
	})

	for item := range observable.Observe() {
		t.Log(item.V)
	}
}
