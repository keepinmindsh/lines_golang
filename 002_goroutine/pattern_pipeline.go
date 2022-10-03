package main

import "fmt"

func PlusOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num + 1
		}
	}()
	return out
}

func ExamplePlusOne() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()

	for num := range PlusOne(PlusOne(c)) {
		fmt.Println(num)
	}
}

type InitPipe func(<-chan int) <-chan int

func Chain(ps ...InitPipe) InitPipe {
	return func(in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(c)
		}
		return c
	}
}

func ExampleWithChain() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()

	for num := range Chain(PlusOne, PlusOne)(c) {
		fmt.Println(num)
	}
}
