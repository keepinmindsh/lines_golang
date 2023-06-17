package sample

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
)

func Test_StringAppendTest(t *testing.T) {

	str := ""
	str1 := "Test Code"

	// 1. += 연산자를 이용해 문자열을 합치는 방법
	start := time.Now()
	for i := 0; i < 100000; i++ {
		str += str1
	}
	elapsed := time.Since(start).Nanoseconds()
	fmt.Printf("+= 연산자를 이용해 문자열을 합치는 방법 - strlen(%d) : %v\n", len(str), elapsed)

	// 2. bytes.Buffers를 이용하 문자열을 합치는 방법
	var b bytes.Buffer
	str = ""
	start = time.Now()
	for i := 0; i < 100000; i++ {
		b.WriteString(str1)
	}
	str = b.String()
	elapsed = time.Since(start).Nanoseconds()
	fmt.Printf("bytes.Buffers를 이용하 문자열을 합치는 방법 - strlen(%d) : %v\n", len(str), elapsed)

	// 3. Sprintf를 이용해 문자열을 합치는 방법
	str = ""
	start = time.Now()
	for i := 0; i < 100000; i++ {
		str = fmt.Sprintf("%s%s", str, str1)
	}
	elapsed = time.Since(start).Nanoseconds()
	fmt.Printf("Sprintf를 이용해 문자열을 합치는 방법 - strlen(%d) : %v\n", len(str), elapsed)

	// 4. Join을 이용해 문자열을 합치는 방법
	str = ""
	mySlice := []string{}
	for i := 0; i < 100000; i++ {
		mySlice = append(mySlice, str1)
	}
	start = time.Now()
	str = strings.Join(mySlice, "")
	elapsed = time.Since(start).Nanoseconds()
	fmt.Printf("Join을 이용해 문자열을 합치는 방법 - strlen(%d) : %v\n", len(str), elapsed)
}
