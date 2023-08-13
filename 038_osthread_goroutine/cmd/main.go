package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())

	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println(runtime.GOMAXPROCS(0))

	s := "Hello, World"

	for i := 0; i < 100; i++ {
		go func(n int) {
			fmt.Println(s, n)
		}(i)
	}

	_, err := fmt.Scanln()
	if err != nil {
		panic(err)
	}
}
