package tests

import (
	"errors"
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_CatchSample(t *testing.T) {

	observable := rxgo.Just(1, 2, errors.New("foo"))().
		OnErrorResumeNext(func(err error) rxgo.Observable {
			return rxgo.Just(3, 4)()
		})

	for item := range observable.Observe() {
		t.Logf("Value : %v", item.V)
	}

}
