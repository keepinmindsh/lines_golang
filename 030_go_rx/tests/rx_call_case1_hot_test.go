package tests

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"sync"
	"testing"
)

func Test_Call_HotObservable(t *testing.T) {

	// 외부 채널로 부터 데이터를 받아오는 방식을 Hot
	ch := make(chan rxgo.Item)

	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go func() {
		for i := 0; i < 10000; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()

	observable := rxgo.FromChannel(ch)

	go func() {
		defer waitGroup.Done()
		for item := range observable.Observe() {
			fmt.Println(fmt.Sprintf("Number: %d", item.V))
		}
	}()

	waitGroup.Wait()
}
