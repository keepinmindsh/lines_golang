package main

import "fmt"

func Example_simpleChannel() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
	}()

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func Example_simpleChannelWithDynamicCount() {
	c := func() <-chan int {
		c := make(chan int)
		go func() {
			defer close(c) // 보내는 쪽에서 close(c)로 채널을 마지막에 닫아주었음.
			c <- 1
			c <- 2
			c <- 3
		}()
		return c
	}()

	for num := range c {
		fmt.Println(num)
	}
}