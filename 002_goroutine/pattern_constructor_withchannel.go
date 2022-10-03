package main

import "fmt"

func FibonacciWithChannel(max int) <-chan int {
	c := make(chan int)

	go func() {
		defer close(c)
		a, b := 0, 1
		for a <= max {
			c <- a
			a, b = b, a+b
		}
	}()
	return c
}

func FibonacciGenerator(max int) func() int {
	next, a, b := 0, 0, 1

	return func() int {
		next, a, b = a, b, a+b
		if next > max {
			return -1
		}
		return next
	}
}

func BabyNames(first, second string) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		for _, f := range first {
			for _, s := range second {
				c <- string(f) + string(s)
			}
		}
	}()
	return c
}

func ExampleBabyNames() {
	for n := range BabyNames("성정명재경", "준호우훈진") {
		fmt.Print(n, ",")
	}

}

func ExampleFibonacci() {

	fmt.Println("FibonacciWithChannel with Generator")

	fibGn := FibonacciGenerator(15)

	for n := fibGn(); n >= 0; n = fibGn() {
		fmt.Println(n)
	}

	fmt.Println("FibonacciWithChannel Normal")

	for fib := range FibonacciWithChannel(15) {
		fmt.Print(fib, ",")
	}

	fmt.Println(" ")
}
