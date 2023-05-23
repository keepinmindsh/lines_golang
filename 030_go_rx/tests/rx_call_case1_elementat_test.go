package tests

import (
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_ElementAt(t *testing.T) {

	observable := rxgo.Just(0, 1, 2, 3, 4)().ElementAt(2)

	for item := range observable.Observe() {
		t.Log(item.V)
	}
}
