package tests

import (
	"context"
	"github.com/reactivex/rxgo/v2"
	"testing"
	"time"
)

func Test_TimerSample(t *testing.T) {
	// FIXME 실제 효과적으로 사용할 수 있는 rxgo.Timer 예제 스터디하기
	observable := rxgo.Timer(rxgo.WithDuration(5 * time.Second)).Map(func(ctx context.Context, i interface{}) (interface{}, error) {
		return time.Now(), nil
	})

	items := observable.Observe()

	for item := range items {
		t.Logf("%v", item.V)
	}
}
