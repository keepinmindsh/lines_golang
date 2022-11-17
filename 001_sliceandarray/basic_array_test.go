package main

import (
	"fmt"
	"testing"
)

func Test_BasicArray(t *testing.T) {
	fruits := [3]string{"사과", "바나나", "토마토"}
	for _, fruit := range fruits {
		fmt.Printf("%s 는 맛있다. \n", fruit)
	}
	// Output :
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
}

func Test_DeclareSlice(t *testing.T) {
	// 기본적으로 빈 슬라이스에는 nil 값이 들어갑니다.
	var fruits []string // 빈문자열 Slice 선언

	fmt.Println(fruits)

	// 슬라이스가 몇개 들어갈지 알고 있는 경우
	apples := make([]string, 100)

	fmt.Println(apples)
}

func Test_CutSlice(t *testing.T) {
	// cut   :  0   1   2   3   4   5
	// values:    1   2   3   4   5
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(nums)
	fmt.Println(nums[1:3])
	fmt.Println(nums[2:])
	fmt.Println(nums[:1])
	fmt.Println(nums[:0])
	fmt.Println(nums[:3])
}
