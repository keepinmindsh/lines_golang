package main

import "fmt"

func main() {
	values := [3]int{1, 2, 3}

	for _, value := range values {
		fmt.Printf("%v 는 숫자다! \n", value)
	}

	fmt.Println("슬라이스 복사 - 첫번째 방법")

	ExampleSliceCopy1()

	fmt.Println("슬라이스 복사 - 두번째 방법")

	ExampleSliceCopy2()

	fmt.Println("슬라이스 복사 - 세번째 방법")

	ExampleSliceCopy3()
}

func ExampleSliceCopy1() {
	src := []int{30, 20, 50, 10, 40}

	dest := make([]int, len(src))
	for i := range src {
		dest[i] = src[i]
	}

	fmt.Println(dest)
}

func ExampleSliceCopy2() {
	src := []int{30, 20, 50, 10, 40}

	dest := make([]int, len(src))

	// 이경우 만약 dest의 사이즈가 적절하지 않다면 복사가 안될 수 있기 때문에 아래의 코드 삽입
	copy(dest, src)

	if n := copy(dest, src); n != len(src) {
		fmt.Println("복사가 덜 됐습니다.")
	}

	fmt.Println(dest)
}

func ExampleSliceCopy3() {
	src := []int{30, 20, 50, 10, 40}

	dest := append([]int(nil), src...)

	for _, value := range dest {
		fmt.Println(value)
	}
}
