package main

import (
	"fmt"
	"sync"
)

func FanIn(ins ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ins))
	for _, in := range ins {
		go func(in <-chan int) {
			defer wg.Done()
			for num := range in {
				out <- num
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func ExampleFanIn() {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	c3 := make(chan int, 10)

	for i := 0; i < 10; i++ {
		c1 <- i
	}
	for i := 0; i < 10; i++ {
		c2 <- i
	}
	for i := 0; i < 10; i++ {
		c3 <- i
	}

	channels := FanIn(c1, c2, c3)

	// todo - 고루틴을 사용하지 않으면 해당 문장에서 멈춰버림!!
	go func() {
		for num := range channels {
			fmt.Print(num)
		}
	}()
}
