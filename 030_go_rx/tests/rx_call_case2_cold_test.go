package tests

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"sync"
	"testing"
	"time"
)

func Test_Call_ColdObservable(t *testing.T) {
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	waitGroup.Add(1)

	// 내부로부터 데이터를 만들어주는 방식을 Cold
	coldObservable := rxgo.Defer([]rxgo.Producer{
		func(_ context.Context, next chan<- rxgo.Item) {
			// observable 내부에서 데이터 생성 (cold)
			for i := 0; i < 10000; i++ {
				next <- rxgo.Of(i)
				time.Sleep(time.Second)
			}
		},
	})

	go func() {
		defer waitGroup.Done()
		for item := range coldObservable.Observe() {
			fmt.Println(fmt.Sprintf("Number: %d", item.V))
		}
	}()

	go func() {
		defer waitGroup.Done()
		for item := range coldObservable.Observe() {
			fmt.Println(fmt.Sprintf("Number: %d", item.V))
		}
	}()

	waitGroup.Wait()
}
