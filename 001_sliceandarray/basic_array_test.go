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

	for _, value := range fruits {
		fmt.Print(`값이 Loop 도나여? #{value}`)
	}

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

	// 제일 마자믹 자리의 수를 뺄 경우, 한개씩 뺄수 있음.
	fmt.Println(nums[:len(nums)-1])
}

// Test_CapacitySlice
/*
=== RUN   Test_CapacitySlice
[1 2 3 4 5]
len: 5
cap: 5

[1 2 3]
len: 3
cap: 5

0
*/
func Test_CapacitySlice(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(nums)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))
	fmt.Println()

	sliced1 := nums[:3]
	fmt.Println(sliced1)
	fmt.Println("len:", len(sliced1))
	fmt.Println("cap:", cap(sliced1))
	fmt.Println()

	var sliceSample []string

	fmt.Println(len(sliceSample))
}

// Test_ArrayForLoop array test for rebinding values
func Test_ArrayForLoop(t *testing.T) {

	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for key, value := range list {

		value = value + 1

		list[key] = value
	}

	for key, value := range list {
		fmt.Printf("%v, %v \r\n", key, value)
	}
}
