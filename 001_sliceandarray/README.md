# Array

```go
package main

import "fmt"

func main() {
	values := [3]int{1, 2, 3}

	for _, value := range values {
		fmt.Println("%s는 숫자다! \n", value)
	}
}
```

# Slice

배열도 쓸모가 있지만 자주 쓰이지 않습니다. 왜냐하면 더 유연한 구조의 슬라이스가 있기 때문입니다. 

```go
package main

func main()  {
	var values []string
}

```

### 슬라이스 초기화 시점에 값을 정의하여 사용하는 방식

```go
package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4, 5}

	for _, value := range values {
		fmt.Println("지금 출력된 숫자 번호는 %s", value)
	}

}
```

### 슬라이스가 몇개인지 알고 있는 경우 편한 방식은 아래와 같이 작성할 수 있습니다. 

```go

fruits := make([]string, n)

```

### 슬라이스 이어붙이기 

```go
package main

func main()  {
	f1 := []int{1,2,3,4,5}
	f2 := []int{1,2,3,4,5}
	f3 := append(f1, f2, ...) // 이어붙이기 
	f4 := append(f1[:2], f2 ...) // 토마토를 제외하고 이어붙이기 
	
}
```

### 슬라이스에서 아이템을 제외시키고 싶은 경우

아이템 값을 제외시키고 싶을 경우, 1 대신에 n을 써서 뒤에서 n 개를 제외할 수 있습니다.
n개 까지 미리 예약해 두었으므로 values = append(values, x)를 한 경우에 n의 값을 넘지 않는다면 복사가 일어나지 않습니다. 

```go

fruitsliced := fruits[:len(fruits) - 1]


```

### 슬라이스 덧붙이기 

```go

fruits = append(fruits, "포도" )

```

### 슬라이스 용량 

슬라이스는 연속된 메모리 공간을 활용하는 것이라서 용량에 제한이 있을 수 밖에 없습니다.  
슬라이스는 길이와 용량의 개념이 있습니다. 
 - 길이 : 실제 아이템을 담고 있는 길이 
 - 용량 : 아이템을 담을 수 있는 용량 

```go

nums := make([]int, 3, 5)

```

그 크기에 딱 맞는 메모리만 받아서 쓰고 있는 것잊요, 공간의 낭비가 없습니다. 
만약 얼마 만큼의 길이가 필요한지 안다면 용량을 사전에 지정하는 것이 성능에 이득일 수 있습니다.  

- 뒤에 얼마나 덧붙일 공간이 있느냐에 따라 용량이 결정됨, 
  - 뒤에 2개를 잘라낸 경우에는 길이는 2만큼 줄어들지만, 기둥 뒤의 공간 있듯이 여전히 2만큼 공간이 뒤에 더 있으므로 용량은 여전히 5가 됩니다.
  - 반면 앞에 2개를 잘내는 경우에는 길이가 2만큼 줄어들고, 뒤에 공간이 없음으로 용량도 3으로 줄어들게 됩니다.

- 중요!

용량을 초과하게 될 경우, 이 경우에는 더 큰 크기의 배열을 새로 하나 더 만들고 슬라이스도 거기에 맞춰 고쳐서 반환합니다.

### 슬라이스의 내부 구현 

**슬라이스는 배열을 가리키고 있는 구조체라고 할 수 있습니다**

```go
package main

func main()  {
	nums := []int{1,2,3,4,5,6}
	
	nums := append(nums, 10)
	
}
```

어떤 경우에라도 nums에 대한 append는 nums로 받아야 합니다. 이 값을 버리게 되면 원래 슬라이스에 길이가 늘어나지 않으므로 
슬라이스는 덧붙이기가 일어나지 않게됩니다. 중요한 것은 **nums에 재할당한 이유는 길이의 증가나 용량의 증가가 결국엔 추가 사이즈가 
필요했던 것인데 nums로 받지 않으면 원하는 덧붙이기 등의 기능이 동작할 수 없다.**

### 슬라이스 복사 

- 첫번째 방법 

```go

func Example_sliceCopy() {
	src := []int{ 30, 20, 50, 10, 40 }

	dest := make([]int, len(src))
	for i := range src {
		dest[i] = src[i]
	}

	fmt.Println(dest)
}

```

- 두번째 방법

```go
package main

import "fmt"

func main() {
	src := []int{30, 20, 50, 10, 40}

	dest := make([]int, len(src))

	// 이경우 만약 dest의 사이즈가 적절하지 않다면 복사가 안될 수 있기 때문에 아래의 코드 삽입
	copy(dest, src)

	if n := copy(dest, src); n != len(src) {
		fmt.Println("복사가 덜 됐습니다.")
	}
}
```

- 세번째 방법

```go
package main

import "fmt"

func main() {
	src := []int{30, 20, 50, 10, 40}

	dest := append([]int(nil), src...)

	for _, value := range dest {
		fmt.Println(value)
	}
}
```

### Slice 삽입

- 첫번째 방법 

```go
package main

func main()  {
	
	a := []int{1,2,3,4,5}
	i := len(a)
	
	x := 10
	
	if i < len(a) {
		a = append(a[:i+1], a[i])
	} else {
		a = append(a, x)
	}
}
```

- 두번째 방법 

```go
package main

func main()  {
	a := []int{1,2,3,4,5}
	x := 10
	
	a = append(a, x)
}
```

### Slice 삭제

```go
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}

	i := int(len(a) / 2)

	a = append(a[:i], a[i+1:]...)

	for _, aValue := range a {
		fmt.Println(aValue)
	}
}
```

삭제시 중요한 점이 하나 더 있는데, 만약에 삭제되는 슬라이스 내부에 포인터가 있는 경우에는 이것이 뒤에 남아 공간에 남아 있으면 가비지 컬렉션이 얼아나지 않기 
때문에 메모리 누수가 일어납니다!   

포인터 슬라이스가 되면 이 영역에 있는 포인터들이 가리키는 객체들까지 메모리 공간을 차지하고 있게 되며, 가비지 컬렉터가 이 객체들을 소멸시킬 수 없습니다.
가비지 컬렉션이 있는 언어에서는 메모리 누수가 일어나지 않는다고 잘못 이해하는 경우가 있는데, 전혀 그렇지 않습니다. 유요한 포인터가 있으면 이것은 가비지 
컬렉션을 할 수 없기 때문에 메모리가 반환되지 않습니다. 이것은 자바 역시 마찬가지 있습니다. 이런 경우에는 해당 포인터를 nil로 해서 삭제해주어야 합니다. 
그러니 뒤에 남아 있는 그림자, 삭제 뒤에 길이 뒤에 있는 공간에 있는 부분을 Nil로 지워주어야 합니다. 슬라이스가 이용하고 있는 배열이 직접 포인터들을 갖고 있는
것이 아니더라도 고조체 배열에서 고조체가 포인터를 갖고 있는 경우에도 동일한 문제가 발생하게 됩니다. 
이럴 때는 구조체 안에 있는 포인터들을 nil로 초기화해주거나 아니면 아예 해당 구조체를 빈 구조체로 덮어쓰거나 해주어야 합니다.

```go
package main

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var pointerAslice *[]int

	pointerAslice = &a

	i := len(*pointerAslice) / 2
	k := i + 3

	copy((*pointerAslice)[i:], (*pointerAslice)[i+k:])

	for i := 0; i < k; i++ {
		// todo - 해당 코드는 이상한데? 포인터 슬라이스에 nil을 넣는 케이스 스터디 필요 
		//pointerAslice[len(*pointerAslice)-1-i] = nil
	}
}
```



