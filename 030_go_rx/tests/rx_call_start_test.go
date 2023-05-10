package tests

import (
	"context"
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_StartSample(t *testing.T) {
	observable := rxgo.Start([]rxgo.Supplier{func(ctx context.Context) rxgo.Item {
		return rxgo.Of(1)
	}, func(ctx context.Context) rxgo.Item {
		return rxgo.Of(2)
	}, func(ctx context.Context) rxgo.Item {
		return rxgo.Of(3)
	}}, rxgo.WithPool(10))

	items := observable.Observe()

	for item := range items {
		t.Logf("%v", item.V)
	}
}
