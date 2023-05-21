package tests

import (
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_Buffer(t *testing.T) {
	observable := rxgo.Just(1, 2, 3, 4, 5, 6, 7)().BufferWithCount(3)

	for item := range observable.Observe() {
		t.Logf("Value : %v", item.V)
	}
}
