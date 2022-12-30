package sample

import "fmt"

func ClosureWithIterable() {
	gen := newEven()
	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(gen())
	gen = nil // release for garbage collection
}

func newEven() func() int {
	n := 0
	// closure captures variable n
	return func() int {
		n += 2
		return n
	}
}
