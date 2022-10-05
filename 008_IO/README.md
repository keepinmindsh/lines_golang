# 입출력 

### io.Reader 

os.Open()은 반환 값이 둘입니다. 하나는 파일 오브젝트 이고 다른 하나는 에러입니다. 이 에러의 값이 nil이 되면 파일을 성공적으로 연 것이 됩니다. 
그러면 이제 f 값을 사용하면 됩니다. 파일을 열지 못한 경우에는 앞으로 진행이 되지 않는 경우가 많으므로 대부분 해당 에러를 반환할 수 있고, 반환이 일어나다가 에러를 
처리할 수 있는 곳에서 다른 방법을 이용하거나 에러로그를 남기거나 아니면 프로그램을 중단 시킬 수 있습니다. 

```go
package main

import (
	"fmt"
	"os"
)

func OpenFileRead() error {

	filename := "IHaveADream.txt"

	file, err := os.Open(filename)

	if err != nil {
		return err
	}
	defer file.Close()

	var num int
	if _, err := fmt.Fscanf(file , "%d\n", &num ); err == nil {
		// 읽은 값 사용 
    }
    
	return nil
}
```

### io.Writer

```go
package main

import (
	"fmt"
	"os"
)

func OpenFileWrite() error {
    num := 100	
	
	filename := "IHaveADream.txt"

	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	if _, err := fmt.Fprintf(file, "%d\n", num); err != nil {
		return err
	}

	return nil
}
```

### Test Read & Write

```go
package main

import (
	"bufio"
	"fmt"
	"io"
)

func WriteTo(writer io.Writer, lines []string) error {
	for _, line := range lines {
		if _, err := fmt.Fprintln(writer, line); err != nil {
			return err
		}
	}
	return nil
}

func ReadFrom(reader io.Reader, lines *[]string) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		*lines = append(*lines, scanner.Text())
    }
	if err := scanner.Err(); err != nil {
		return err
    }
	return nil
}
```

