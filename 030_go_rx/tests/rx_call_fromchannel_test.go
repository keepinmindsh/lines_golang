package tests

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Test_FromChannel(t *testing.T) {
	ch := make(chan rxgo.Item)

	observable := rxgo.FromChannel(ch)

	go func() {
		defer close(ch)
		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(i)
		}
	}()

	observable.DoOnNext(func(i interface{}) {
		fmt.Println(fmt.Sprintf("%s", i))
	})

	subscribe(observable)

	// this is won't receive data, because this is already used from channel
	subscribe(observable)

}

func subscribe(observable rxgo.Observable) {
	subscriber := observable.Observe()

	for item := range subscriber {
		fmt.Println(fmt.Sprintf("%s", item.V))
	}
}
