package tests

import (
	"errors"
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_Catch_OnErrorReturnItem_Sample(t *testing.T) {
	observable := rxgo.Just(1, errors.New("2"), 3, errors.New("4"), 5, 6)().
		OnErrorReturnItem("foo")

	for item := range observable.Observe() {
		t.Logf("Value : %v", item.V)
	}
}
