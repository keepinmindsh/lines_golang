package tests

import (
	"context"
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_CreateSample(t *testing.T) {
	observable := rxgo.Create([]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
		next <- rxgo.Of(1)
		next <- rxgo.Of(2)
		next <- rxgo.Of(3)
	}})

	observable.DoOnNext(func(i interface{}) {
		t.Logf("%v", i)
	})

	items := observable.Observe()

	for item := range items {
		t.Logf("%v", item.V)
	}
}
