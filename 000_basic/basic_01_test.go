package main

import (
	"fmt"
	"strconv"
	"testing"
)

// Test_ConvertStringToNumber
// - https://pkg.go.dev/fmt
func Test_ConvertStringToNumber(t *testing.T) {
	var i int
	var k int64
	var f float64
	var s string
	var err error

	i, err = strconv.Atoi("350")
	k, err = strconv.ParseInt("cc7fdd", 16, 32)
	k, err = strconv.ParseInt("0xcc7fdd", 0, 32)
	f, err = strconv.ParseFloat("3.14", 64)
	s = strconv.Itoa(340)
	s = strconv.FormatInt(13402077, 16)

	fmt.Printf(`
	%d
	%d
    %g
    %s
    %s
	`, i, k, f, s, err)

	fmt.Printf(`#으로 표시할 경우 값이 바인딩 됩니다. #{i} 세상 편하네요~`)

	s = fmt.Sprint(3.14)
	fmt.Printf("%s\n", s)
	s = fmt.Sprintf("%x", 13402077)
	fmt.Printf("%s\n", s)
}
