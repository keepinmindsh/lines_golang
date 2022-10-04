# Map

Go 언어에서 map은 해시 테이블로 구현됩니다. 해시맵은 키와 값으로 되어 있는데 키를 이용해서 값을 상수시간에 가져올 수 있습니다.  
그 대신에 해시맵에는 순서가 없습니다. 

```go
package main 

var m map[KeyType]valueType

```

```go
package main

import "fmt"

func main() {
	m := make(map[string]int)

	fmt.Println(m)
}
```

```go
package main

import "fmt"

func main() {
	m := map[string]int{}

	fmt.Println(m)
}
```

맵에서 읽을 때는 두가지 방법이 있습니다. m[key]를 이용하면 맵의 값을 읽을 수 있습니다. 만일 해당 키가 없으면 값의 자료형의 기본 값을 반환합니다. 
맵을 읽을 때 두 개의 변수로 받게 되면, 두 번째 변수에 키가 존재하는지 여부를 bool 형으로 받을 수 있습니다. 

```go
package main

import "fmt"

func main() {
	m := map[string]int{}

	value, ok := m["value"]
	
	fmt.Printf("value : %s %s", value, ok)
}
```

## Map 사용하기 

슬라이스와 다른 점은 맵을 이용할 때에는 맵 변수 자체에 다시 할당하지 않으므로 포인터를 취하지 않아도 맵을 변경할 수 있습니다. 

```go
package main

func Count(s string, codeCount map[rune]int) {
	for _, r := range s {
		codeCount[r]++
	}
}

```

물론 맵을 다른 맵으로 바꿔치기 하고 싶으면 포인터를 넘겨야 합니다만 그런 경우는 흔치 않습니다. 반면에 슬라이스는 추가할 때 a = append(a, ...) 와 같은 형식이 사용되므로 포인터를 넘겨야 추가할 수 있습니다.

```go
package main

import "fmt"

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
```