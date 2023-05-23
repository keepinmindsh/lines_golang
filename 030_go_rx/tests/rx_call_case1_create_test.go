package tests

import (
	"context"
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_Create(t *testing.T) {
	observable := rxgo.Create([]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
		next <- rxgo.Of(1)
		next <- rxgo.Of(2)
		next <- rxgo.Of(3)
	}})

	for item := range observable.Observe() {
		t.Log(item.V)
	}
}
