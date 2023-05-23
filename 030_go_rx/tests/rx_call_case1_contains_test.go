package tests

import (
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_ContainTest(t *testing.T) {
	observable := rxgo.Just(1, 2, 3, 4, 5, 6)().Contains(func(i interface{}) bool {
		t.Logf("Test : %v", i.(int))
		return i == 3
	})

	for item := range observable.Observe() {
		t.Log(item.V)
	}
}
