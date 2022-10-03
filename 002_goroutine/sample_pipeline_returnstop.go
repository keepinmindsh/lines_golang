package main

import (
	"fmt"
	"runtime"
	"time"
)

func ReturnStopPipeline() {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 3; i < 103; i += 10 {
			c <- i
		}
	}()

	nums := PlusOne(PlusOne(PlusOne(PlusOne(c))))

	for num := range nums {
		fmt.Println(num)
		if num == 18 {
			break
		}
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())

	for _ = range nums {
		// Consume all nums
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
}

func PlusOneWithChannel(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int) // 양방향으로 채널이 사용될 수 있음
	go func() {
		defer close(out)
		for num := range in {
			select {
			case out <- num + 1:
			case <-done:
				return
			}
		}
	}()
	return out
}

func ReturnStopPipelineWithChannel() {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 3; i < 103; i += 10 {
			c <- i
		}
	}()
	done := make(chan struct{})
	nums := PlusOneWithChannel(done, PlusOneWithChannel(done, PlusOneWithChannel(done, PlusOneWithChannel(done, PlusOneWithChannel(done, c)))))
	for num := range nums {
		fmt.Println(num)
		if num == 18 {
			break
		}
	}
	close(done)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
	for _ = range nums {
		// Consume All nums
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
}
