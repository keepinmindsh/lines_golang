package tests

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"testing"
	"time"
)

func Test_IntervalCallSample(t *testing.T) {

	observable := rxgo.Interval(rxgo.WithDuration(5 * time.Second))

	observable.DoOnNext(func(i interface{}) {
		t.Logf("%v", i)
	})

	observe := observable.Observe()

	item := <-observe

	fmt.Printf("%v", item.V)
}
