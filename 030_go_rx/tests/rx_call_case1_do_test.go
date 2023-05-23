package tests

import (
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_DoSample(t *testing.T) {
	<-rxgo.Just(1, 2, 3)().
		DoOnNext(func(i interface{}) {
			t.Log(i)
		})
}
