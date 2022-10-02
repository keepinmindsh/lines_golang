package main

import (
	"fmt"
	"sync"
)

func Distribute(p InitPipe, n int) InitPipe {
	return func(ints <-chan int) <-chan int {
		cs := make([]<-chan int, n)
		for i := 0; i < n; i++ {
			cs[i] = p(ints)
		}
		return FanIn(cs...)
	}
}

func ExampleDistribute() {
	fmt.Println("------------ Start ExampleDistribute ------------")
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()

	var wg sync.WaitGroup

	out := Chain(PlusOne, Distribute(Chain(PlusOne, PlusOne, PlusOne), 10), PlusOne)(c)

	wg.Add(1)

	go func() {
		for num := range out {
			fmt.Println(num)
		}

		wg.Done()
	}()

	wg.Wait()

	fmt.Println("------------ End ExampleDistribute ------------")
}
