package tests

import (
	"context"
	"errors"
	"github.com/cenkalti/backoff"
	"github.com/reactivex/rxgo/v2"
	"testing"
	"time"
)

func Test_RetrySample(t *testing.T) {
	backOffCfg := backoff.NewExponentialBackOff()
	backOffCfg.InitialInterval = 10 * time.Millisecond

	observable := rxgo.Defer([]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
		next <- rxgo.Of(1)
		next <- rxgo.Of(2)
		next <- rxgo.Error(errors.New("foo"))
	}}).BackOffRetry(backoff.WithMaxRetries(backOffCfg, 2))

	items := observable.Observe()

	for item := range items {
		t.Logf("Value : %v", item.V)
		t.Logf("Error : %v", item.E)
	}
}
