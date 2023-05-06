package tests

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"testing"
	"time"
)

func Test_ConnectableObservable(t *testing.T) {
	ch := make(chan rxgo.Item)
	go func() {
		ch <- rxgo.Of(1)
		ch <- rxgo.Of(2)
		ch <- rxgo.Of(3)
		close(ch)
	}()
	// Create a Connectable Observable
	observable := rxgo.FromChannel(ch, rxgo.WithPublishStrategy())

	// Create the first Observer
	observable.DoOnNext(func(i interface{}) {
		fmt.Printf("First observer: %d\n", i)
	})

	// Create the second Observer
	observable.DoOnNext(func(i interface{}) {
		fmt.Printf("Second observer: %d\n", i)
	})

	disposed, cancel := observable.Connect(context.Background())
	go func() {
		// Do something
		time.Sleep(time.Second)
		// Then cancel the subscription
		cancel()
	}()

	// Wait for the subscription to be disposed
	<-disposed.Done()
}
