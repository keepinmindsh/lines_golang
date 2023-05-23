package tests

import (
	"context"
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_DistinctUntilChanged(t *testing.T) {
	observable := rxgo.Just(1, 2, 2, 1, 1, 3)().
		DistinctUntilChanged(func(ctx context.Context, i interface{}) (interface{}, error) {
			return i, nil
		})

	for item := range observable.Observe() {
		t.Log(item.V)
	}
}
