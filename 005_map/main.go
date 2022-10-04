package main

import "fmt"

func main() {
	m := map[string]int{}

	m["value"] = 100

	fmt.Printf("Value : %d", m["value"])

	Count("가나다나", map[rune]int{'가': 1, '나': 2, '다': 1})

	ExampleCount()
}
