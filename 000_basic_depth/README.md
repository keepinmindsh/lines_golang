# Golang 을 깊이 있게 알아볼까? 

## 바로가기 

### [문자열 이어붙이기](#문자열-이어붙이기)

---

## 문자열 이어붙이기 

```shell
# 최종 결과 
=== RUN   Test_StringAppendTest
+= 연산자를 이용해 문자열을 합치는 방법 - strlen(900000) : 2379301291
bytes.Buffers를 이용하 문자열을 합치는 방법 - strlen(900000) : 467875
Sprintf를 이용해 문자열을 합치는 방법 - strlen(900000) : 5471952125
Join을 이용해 문자열을 합치는 방법 - strlen(900000) : 888042
--- PASS: Test_StringAppendTest (7.85s)
```

**실질적으로 가장 빠른 방식**은 아래의 코드예제와 유사하다. 

````go
package sample 

import (
	"bytes"
	
	"fmt"
	"time"
	"testing"
)

func Test_ByteBuffers(t *testing.T) {
	str := ""
	str1 := "Test Code"
	start := time.Now()
	
	// 2. bytes.Buffers를 이용하 문자열을 합치는 방법
	var b bytes.Buffer
	str = ""
	start = time.Now()
	for i := 0; i < 100000; i++ {
		b.WriteString(str1)
	}
	str = b.String()
	elapsed := time.Since(start).Nanoseconds()
	fmt.Printf("strlen(%d) : %v\n", len(str), elapsed)
}
````

WriteString() 함수는 결국 Buffer 구조가 가지고 있는 byte.slice(buf)에 인자로 받는 string을 이어붙이기하는 함수이다. 
그런데, 처음에 이미 할당된 slice 공간을 이용하려고 시도하지만, 공간이 모자라면 결국 slice의 capacity를 늘린후 (grow) string 을 복사하게 된다. 
따라서, WriteString() 함수는 slice에 계속해서 데이터를 추가해 나가는 작업의 속도다! 


> 참조 [Go – String 을 어떻게 빠르게 이어붙일까?(String Concatenation)](http://cloudrain21.com/go-how-to-concatenate-strings)