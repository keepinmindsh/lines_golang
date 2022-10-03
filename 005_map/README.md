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