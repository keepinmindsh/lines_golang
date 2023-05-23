package tests

import (
	"context"
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_DistinctTest(t *testing.T) {
	observable := rxgo.Just(1, 2, 2, 3, 4, 4, 5)().
		Distinct(func(ctx context.Context, i interface{}) (interface{}, error) {
			return i, nil
		})

	for item := range observable.Observe() {
		t.Log(item.V)
	}
}
