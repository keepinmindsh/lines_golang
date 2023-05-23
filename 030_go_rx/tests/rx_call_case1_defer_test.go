package tests

import (
	"context"
	"github.com/reactivex/rxgo/v2"
	"testing"
	"time"
)

func Test_DeferSample(t *testing.T) {
	observable := rxgo.Defer([]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
		next <- rxgo.Of(1)
		next <- rxgo.Of(2)
		next <- rxgo.Of(3)
	}})

	for item := range observable.Observe() {
		time.Sleep(1 * time.Second)
		t.Log(item.V)
	}
}
