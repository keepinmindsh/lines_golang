package main

import "fmt"

func main() {
	values := [3]int{1, 2, 3}

	for _, value := range values {
		fmt.Println("%s는 숫자다! \n", value)
	}
}
