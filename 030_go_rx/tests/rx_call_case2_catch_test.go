package tests

import (
	"errors"
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_Catch_OnErrorReturn_Sample(t *testing.T) {

	observable := rxgo.Just(1, 2, errors.New("3"), 4, errors.New("5"), 6)().
		OnErrorReturn(func(err error) interface{} {
			return err.Error()
		})

	for item := range observable.Observe() {
		t.Logf("Value : %v", item.V)
	}
}
