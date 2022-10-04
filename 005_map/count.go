package main

import (
	"fmt"
	"sort"
)

func Count(s string, codeCount map[rune]int) {
	for _, r := range s {
		codeCount[r]++
	}
}

func ExampleCount() {
	codeCount := map[rune]int{}
	Count("가나다나", codeCount)
	var keys sort.IntSlice
	for key := range codeCount {
		keys = append(keys, int(key))
	}
	for _, key := range keys {
		fmt.Println(string(key), codeCount[rune(key)])
	}
}
