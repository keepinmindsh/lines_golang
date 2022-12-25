# Type Assertion / Type Conversion 

### Type Converison 

- 형 변환이다. golang 에서는 형변환을 하기 위해서는 아래와 같이 명시적으로 해줘야 한다.

```go
package main

import "fmt"

func main() {
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Printf("%v, %v, %v", i, f, u )
}

```

### Type Assertion 

- Type Assertion 은 interface type의 value의 타입을 확인하는 것이다.

인터페이스가 가지고 있는 실제 값에 접근할 수 있게 해준다. 

```go
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TypeAssertion(t *testing.T) {
	var i interface{} = "hello"

	v, ok := i.(int)

	assert.Equal(t, true, ok)
	assert.Equal(t, "hello", v)
}
```

```go
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TypeAssertion_Sample(t *testing.T) {
	var i interface{} = "hello"

	s := i.(string)
	assert.Equal(t, "hello", s)

	s, ok := i.(string)
	assert.Equal(t, true, ok)

	f, ok := i.(float64)
	assert.Equal(t, true, ok)
	assert.Equal(t, "hello", f)

	/*
		panic: interface conversion: interface {} is string, not float64 [recovered]
			panic: interface conversion: interface {} is string, not float64
	*/
	f = i.(float64)
	assert.Equal(t, "hello", f)
}

```