package tests

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"sync"
	"testing"
)

func Test_Call_ObservableFromChannel(t *testing.T) {
	ch := make(chan rxgo.Item)
	observable := rxgo.FromEventSource(ch)

	var wg sync.WaitGroup

	wg.Add(1)
	wg.Add(1)
	go func() {
		defer close(ch)
		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(i)
		}
	}()

	observable.DoOnNext(func(i interface{}) {
		fmt.Println(fmt.Sprintf("test1 - %d", i))
	})

	go func() {
		defer wg.Done()
		subscribeForFromChannel(observable, "sub1")
	}()

	go func() {
		defer wg.Done()
		subscribeForFromChannel(observable, "sub2")
	}()

	wg.Wait()
}

func subscribeForFromChannel(observable rxgo.Observable, subId string) {
	subscriber := observable.Observe()

	for item := range subscriber {
		fmt.Println(fmt.Sprintf("test2 - %d , %s", item.V, subId))
	}
}
