# Go Study

### Package Import

- Package 를 import 했을 때, 외부에 공개 가능한 함수는 영문자로 첫글자가 반드시 대문자여야한다.

### Go - Pointer 를 쓰는 이유

- Stack & Heap
    - 실행 시 동적으로 메모리를 확보하는 영역으로 스택과 힙이 있다.
    - 스택 메모리는 함수 호출 스택을 저장하고 로컬 변수, 인수, 반환 값도 여기에 둔다.
    - 스택의 Push와 Pop은 고속이므로 객체를 스택 메모리에 저장하는 비용은 작다. 단 함수를 나오면 스택이 Pop 되어 해제되므로 함수의 수명을 넘는 객체는 살 수 없다.
    - 힙 메모리는 콜 스택과는 관계 없으므로 함수 범위에 얽매이지 않고 객체를 저장해 둔다. 다만 빈 영역을 찾고, GC로 쓸모 없게된 객체를 회수하기도 하므로 처리 비용이 든다.
    - Go 언어는 컴파일러가 객체를 스택에 확보할지 힙에 확보할지 결정하므로 프로그래머가 의식할 필요는 보통 없다.

- 이를 정리해보면
    - 함수 내에서만 사용되는 값은 스택에 둔다.
    - 어떤 함수 내에서 확보한 값이 함수 밖에서도 필요하게 된다면 힙에 놓인다

##### Compiler 에 Flag를 전달하는 방법

```shell
> go build -gcflags -m hello.go
```

##### Stack

```go

func test1() {
    var d Duck = Duck{}
    d.Sound()
}

func test2() {
    var d *Duck = &Duck{}
    d.Sound()
}

func test3() Duck {
    var d Duck = Duck{}
    return d
}

func test4() Duck {
    var d *Duck = &Duck{}
    return *d
}

```

##### Heap

```go

func test5() *Duck {
    var d Duck = Duck{}
    return &d
}

func test8() Sounder {
  var d Sounder = &Duck{}
  return d
}

```

###### 정리

- 함수 내에서만 사용되는 값은 스택
- 함수 밖에서도 값이 필요하게 된다면 힙
- new 해도 스택인 경우가 있다
- 함수 내에서만 사용되는 값이면
- 로컬 변수의 주소를 반환해도 된다
- 힙에 두도록 컴파일러가 다룬다
- 인라인 전개에 따라 잘 최적화해 준다
- interface에 대입하면 힙

> [참조 문서 - Stack/Heap](https://jacking75.github.io/go_stackheap/)

###### 포인터를 쓰는 이유

포인터는 타입이다. 대신 포인터가 가리키는 변수의 메모리 주소를 갖는다. 즉 포인터 변수를 통해 다른 변수의 메모리 주소를 참조해 무언가 가능할 것 같은 느낌이다.

- 메모리 주소
- 역참조

```go

// TestPointer2 포인터의 역참조를 통해 메모리 주소로 value값 접근이 가능할 수 있었다.
func TestPointer2() {
	var example int
	var pointer *int
	example = 3

	pointer = &example

	fmt.Println(&example)
	
	// 포인터의 역참조 (*포인터변수명 = 대입값)
	*pointer = 5

	fmt.Println(pointer)

	fmt.Println(example)
}

```

go언어는 기본적으로 값에 의한 호출(Call by value)이므로 매개변수를 복사해서 함수 내부로 전달하는데 포인터를 통해 본래의 값을 변경할 수 있다.

### go 에서의 인터페이스

go 에서 인터페이스는 Java와 달리 구현의 개념이 아닌 Interface가 실제 함수의 추상을 가지고 이를 사용할 수 있는 구조로 되어 있음

```go


func (p Pet) getName() {
	fmt.Printf("pet name is %s.\n", p.Name)
}

func (p Pet) getAge() {
	fmt.Printf("pet age is %d.\n", p.Age)
}

type GetName interface {
	getName()
}

type GetAge interface {
	getAge()
}

type PetInfo interface {
	GetName
	GetAge
}

func interfaceSample() {
	var interfaceSample PetInfo

	interfaceSample.getName()
}


```

### Go Method Receiver

Go는 클래스가 없다. 하지만 특정 타입에 대한 메소드를 만들 수 있다.

Vertex와 Vertex3D 구조체가 있다.

- Vertex는 Abs라는 메소드를 갖는다.

```go
package main

import "math"

type Vertex struct {
  X, Y float64
}

func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
```
- Vertex3D는 Abs 라는 함수를 통해 연산이 가능하다.

Go에서는 타언어의 인스턴스와 비교해보면 필요한 구조체만 해당 함수에 접근하면 되니까 메모리 다이어트가 필요 없게 되는 것 같다.

##### Pointer Receiver

- 메소드가 receiver pointer의 값을 수정할 수 있습니다.
- 메소드 호출에 따른 값의 복사를 방지하기 위해서입니다. 구조체가 클 수록 효율이 좋습나다.


### unsupported Scan, storing driver.Value type []uint8 into type *time.Time 발생시

Alright, I found the solution, thanks this answer. The problem goes by adding ?parseTime=true to the db mapper. Like this:

```go
package main

func main()  {
  db, err := sql.Connect("mysql", "myuser:mypass@tcp(127.0.0.1:3306)/mydb?parseTime=true")
}

```

### 배역과 슬라이스의 차이점

- 배열의 크기는 고정, 슬라이스의 크기는 가변
- 배열은 복사함녀 별도 메모리를 생성, 슬라이스는 본사할 경우 같은 곳을 참조

##### 배열 선언 방법

```go
package main

import "fmt"

func main() {
  arr1 := [5]int{1, 2, 3, 4, 5}
  arr2 := [...]int{6, 7, 8, 9, 10}

  fmt.Println(arr1)
  fmt.Println(arr2)
}

```

##### 슬라이스 선언 방법

```go
package main

import "fmt"

func main() {
  slice1 := []int{1, 2, 3, 4, 5}

  fmt.Println(slice1)
}

```

##### 배열과 슬라이스의 크기 변경

- 배열의 초기 선언된 고정 값에서 크기를 늘릴 수 없기 때문에 에러가 발생함.

```go
package main

import "fmt"

func main() {
  arr1 := [5]int{1, 2, 3, 4, 5}
  arr1 = append(arr1, 6)

  fmt.Println(arr1)
}

```

- 슬라이스의 초기에 선언된 크기에서 아이템이 추가될 경우 동적으로 아이템이 추가되었습니다.

```go
package main

import "fmt"

func main() {
  slice1 := []int{1, 2, 3, 4, 5}
  slice1 = append(slice1, 6)

  fmt.Println(slice1)
}
```

##### 배열과 슬라이스의 복사

배열의 경우에는 변수를 다른 변수에 복사하면 같은 값을 가지는 새로운 배열이 생깁니다.
그리고 원래 있던 배열과 복사한 배열은 메모리 상에 각각의 위치를 가지게 됩니다.

- 배열 복사 확인

```go
package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := a

	// 메모리 위치 확인
	fmt.Println(&a[0]) // 0xc8200122e0
	fmt.Println(&b[0]) // 0xc820012300

	// 복사한 배열 값을 변경
	b[0] = 0
	fmt.Println(a) // [1 2 3]
	fmt.Println(b) // [0 2 3]
}
```

- 슬라이스 복사 확인

```go
package main

import "fmt"

func main() {
	s1 := []int{1,2,3}
	s2 := s1

	// 메모리 위치 확인
	fmt.Println(&s1[0]) // 0xc8200122e0
	fmt.Println(&s2[0]) // 0xc8200122e0

	// 복사한 슬라이스 값을 변경
	s2[1] = 0

	// 같은 곳을 잠초하고 있기때문에 원본도 변경
	fmt.Println(s1) // [1 0 3]
	fmt.Println(s2) // [1 0 3]
}
```

##### 슬라이스를 복사하는 경우

copy 함수를 사용해 슬라이스를 복사해 사용하면 어느 한쪽 값을 변경해도 서로 영향이 없습니다.
슬라이스를 복사해야 하는 경우에는 단순 대입이 아닌 copy 함수를 사용해서 만들어야 합니다.

```go
package main
import "fmt"

func main() {
	slice := []int{0, 10, 20, 30}

	// 새로운 슬라이스를 생성
	copyslice := make([]int, len(slice))
	// copy 함수로 슬라이스 복사
	copy(copyslice, slice)

	// 값 변경
	copyslice[0] = 100

	fmt.Println(copyslice) // [100 10 20 30]
	fmt.Println(slice)     // [0 10 20 30]
}
```

### 전역 변수를 사용하지 않도록 해야한다.

router 상에 handler 함수 선언시 DB 설정 등에 대해서 함수에 정의할 수 없으므로 전역변수를 사용해야하는 이슈가 생길 수 있는데,
이를 경우에는 구조체를 이용하여 전역 변수 대신 구조체 내의 변수로 사용해야 한다.


```go
package main

import (
  "database/sql"
  "fmt"
  "log"
  "net/http"
)


type HelloHandler struct {
	db *sql.DB
}

func (h *HelloHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var name string
	// Execute the query.
	row := h.db.QueryRow("SELECT myname FROM mytable")
	if err := row.Scan(&name); err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	// Write it back to the client.
	fmt.Fprintf(writer, "hi %s!\n", name)
}

func HttpHandlerStart() {
	// Open our database connection.
	db, err := sql.Open("postgres", "")
	if err != nil {
		log.Fatal(err)
	}
	// Register our handler.
	http.Handle("/hello", &HelloHandler{db: db})
	http.ListenAndServe(":8080", nil)

}

```