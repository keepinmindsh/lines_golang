# Interface 

Golang 에서는 타입 선업 키워드(type)을 사용하여 인터페이스를 선언하며, 인터페이스 명 뒤에 인터페이스 선언 키워드를 추가하여 인터페이스를 정의합니다.   s
인터페이스 안에는 구현이 없는 메서드(메서드명, 파라메터, 리턴 타입만 선언)를 선언하며, 이렇게 선언된 메서드들을 가지고 있는 타입을 우리가 정의한 인터페이스로 
인식하겠다는 것을 의미합니다. Golang에서는 인터페이스도 하나의 타입이며, 인터페이스로 변수를 선언할 수도 있습니다. 

```go
package main 

import "fmt"

type SampleInterface interface {
	SampleMethod()
}

func main() {
	var s SampleInterface
	fmt.Println(s)
}

```

### 인터페이스 규칙 

- 메서드는 반드시 메서드 명이 있어야 한다. 
- 매개 변수와 반환이 다르더라도 이름이 같은 메서드는 있을 수 없다. 
- 인터페이스 에서는 메소드 구현을 포함하지 않는다.

# 덕 타이핑 

덕 타이핑 - 어떤 객체가 어떤변수를 가지고 있고, 어떤 함수를 가지고 있는지와 관계없이, 해당 객체를 사용하는 쪽에서, 이런 함수를 가지고 있다면 
이런 타입으로 보겠다고 정의할 수 있다.

> [Interface](https://dev-yakuza.posstree.com/ko/golang/interface/)