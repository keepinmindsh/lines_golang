package main

import "fmt"

func main() {
	// 알파벳을 10진수로 표현함.
	fmt.Println("Hello World"[0])

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false) // 하나라도 참이면 참!
	fmt.Println(!true)         // NOT 조건
}
