# Slice

슬라이스가 몇개인지 알고 있는 경우 편한 방식 - 아래 코드 참조 

```go

fruits := make([]string, n)

```

아이템 값을 제외시코 싶을 경우, 1 대신에 n을 써서 뒤에서 n 개를 제외할 수 있습니다.


```go

fruitsliced := fruits[:len(fruits) - 1]


```

슬라이스 덧붙이기 


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

만약 얼만만큼의 길이가 필요한지 안다면 용량을 사전에 지정하는 것이 성능에 이득일 수 있습니다.  


- 중요!

용량을 초과하게 될 경우, 이 경우에는 더 큰 크기의 배열을 새로 하나 더 만들고 슬라이스도 거기에 맞춰 고쳐서 반환한다. 

### 슬라이스 복사 

func Example_sliceCopy() {
	src := []int{ 30, 20, 50, 10, 40 }

	dest := make([]int, len(src))
	for i := range src {
		dest[i] = src[i]
	}

	fmt.Println(dest)
}

