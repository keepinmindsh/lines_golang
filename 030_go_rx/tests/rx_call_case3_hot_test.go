package tests

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"sync"
	"testing"
)

func Test_Call_HotObservable_MultipleObservable(t *testing.T) {
	ctx := context.Background()

	// 외부 채널로 부터 데이터를 받아오는 방식을 Hot
	ch := make(chan rxgo.Item)

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()

	// connectable observable 생성
	observable := rxgo.FromChannel(ch, rxgo.WithPublishStrategy())

	go func() {
		defer waitGroup.Done()

		// DoOnNext 메소드를 통해야 다중 연결이 된다.
		observable.DoOnNext(func(i interface{}) {
			fmt.Println(fmt.Sprintf("Step1 Number: %d", i.(int)-1))

		})

		observable.DoOnNext(func(i interface{}) {
			fmt.Println(fmt.Sprintf("Step2 Number: %d", i.(int)-1))
		})

		observable.Connect(ctx)

		for item := range observable.Observe() {
			fmt.Println(fmt.Sprintf("Final Number: %d", item.V))
		}
	}()

	waitGroup.Wait()
}
