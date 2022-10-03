package main

import (
	"fmt"
	"sync"
)

func FindMinimumValue() {
	fmt.Println(parallelMin(
		[]int{
			83, 46, 49, 23, 97, 12, 11, 46, 49, 27, 5,
		}, 2))
}

func parallelMin(a []int, n int) int {
	if len(a) < n {
		return min(a)
	}

	mins := make([]int, n)
	size := (len(a) + n - 1) / n
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			begin, end := i*size, (i+1)*size
			if end > len(a) {
				end = len(a)
			}
			mins[i] = min(a[begin:end])
		}(i)
	}
	wg.Wait()
	return min(mins)
}

func min(a []int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]

	for _, e := range a[1:] {
		if min > e {
			min = e
		}
	}
	return min
}
