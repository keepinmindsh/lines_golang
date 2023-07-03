# Goroutine

go를 붙여서 함수를 호출하게 되면 goroutinetest 호출과 현재 함수의 흐름은 메모리를 고유하는 논리적으로 별개의 흐름이됩니다.  
여기서 논리적으로 별개의 흐름이라고 한 이유는 물리적으로 별개의 흐름이 되는 것과는 구분 되기 때문입니다. 

```go
package main

import "fmt"

func main() {
	fmt.Println("I can get anything if I try to get it with all my effort")
    go goroutinetest(1,2,3)
	fmt.Println("All things happen whether you want or not")
}

func goroutinetest(x, y, z int) {
	fmt.Println("%s, %s, %s", x, y, z)
}
```

## 병렬성과 병행성 

커피를 마시면서 신문을 보고 있는 사람이 있다면 물리적으로 두 흐름이 동시에 수행되는 것은 아닙니다. 커피를 마시기 위하여 신문을 보는 것을 짧은 시간 
잠시 중단하고 커피를 한 모금 마신뒤에 다시 신문 보는 일로 돌아오는 경우는 물리적으로 두 흐름이 있지는 않지만, 동시에 두 가지를 하고는 있습니다. 
이를 동시성 혹은 병행성이라고 합니다. 물리적으로 동시에 수행되지는 않지만 순차적으로 수행되는 것과는 큰 차이가 있습니다. 즉, 커피 한잔을 모두 다 마시는
동안 신문을 전혀 보지 않고, 반드시 커피를 다 마시고 난 다음에 신문을 보기 시작해야 한다는 것이 아닙니다. 이 과정에서 신문 기사의 다섯번째 문단을 읽는 
것 중에서 어느 것을 먼저 해야하는지는 그다지 상관 없는 것이 됩니다.  

## 동시성과 병렬성은 다르지만, 

동시성이 있어야 병렬성이 생기게 됩니다. 서로 어느것이 먼저 되어야하는 의존 관계가 있는 것은 함께 진행될 수 없지요. 장갑을 끼면서 동시에 구두를 신을 수는 있겠지만, 
양말을 신으면서 동시에 구두를 신을 수는 없습니다. 양말을 신고 나야 구두를 신을 수 있으니까요, 이 둘 사이에 동시성이 없으므로 병렬성이 생기지 않는 것입니다. 
즉, 장값을 먼저 끼거나 구두를 먼저 신는 것은 어느 것을 먼저해도 상관이 없지만, 양말은 구두를 신기전에 신어야 합니다. 

## Main Goroutine는 내부의 서브 Goroutine의 처리 시간을 보장하지 않습니다.

아래의 코드에서 main 함수 내부의 sub goroutine의 처리시간을 보장하지 않습니다. main goroutine이 종료되면 서브는 강제적으로 종료됩니다.

```go
package main

import "fmt"

func main() {
	go func() {
		fmt.Println("In goroutine")
	}()
	fmt.Println("In main routine")
}
```

**그러면 어찌해야 할까요?**

## 고루틴 기다리기 

실제 코드로 한번 볼까요?

```go
package main 

func main() {
	urls := []string{
		"http://image.com/img01.jpg",
		"http://image.com/img02.jpg",
		"http://image.com/img03.jpg",
	}

	for _, url := range urls {
		go func(url string) {
			if _, err := Download(url); err != nil {
				log.Fatal(err)
			}
		}(url)
	}

	filenames, err := filepath.Glob("*.jpg")

	if err != nil {
		log.Fatal(err)
	}

	err = WriteZip("images.zip", filenames)

	if err != nil {
		log.Fatal(err)
	}
}
```

위의 프로그램을 실행하면 어떻게 될까요? 원하는 대로 압축이 될 까요? ...?  
파일이 다운로드 되기도 전에 압축하려고 들 수 있기 때문에 압축이 되지 않습니다. 저 위의 코드에 따르면 각각의 파일을 다운로드 하는 것과 압축하는 것 중 
어느 것을 먼저 수행하여도 좋은 동시성있는 작업들입니다. 그러나 실제로는 그렇지 않습니다. 파일 다운로드가 완료되지 않으면 압축도 할수 없습니다.  
즉, 압축 작업은 파일들이 모두 다운로드될 때까지 기다린후 수행되야 되겠습니다. 즉, 압축 작업은 파일들이 모두 다운로드 될때까지 기다링 후 수행해야 되겠습니다. 
이럴 때 이용할 수 있는 것이 sync.WaitGroup입니다.

```go
package main

import "sync"

func main() {

	var wg sync.WaitGroup
	
	urls := []string{
		"http://image.com/img01.jpg",
		"http://image.com/img02.jpg",
		"http://image.com/img03.jpg",
	}

	wg.Add(len(urls))

	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			if _, err := Download(url); err != nil {
				log.Fatal(err)
			}
		}(url)
	}
	
	wg.Wait()

	filenames, err := filepath.Glob("*.jpg")

	if err != nil {
		log.Fatal(err)
	}

	err = WriteZip("images.zip", filenames)

	if err != nil {
		log.Fatal(err)
	}
}
```

wg 에는 기본 값이 0으로 맞춰져 있는 카운터가 들어가 있습니다. Wait() 함수는 이 카운터가 0이 될때까지 기다립니다. Add() 함수는 호출될 때 마다 숫자를 더합니다. 
wg.Done()은 사실상 wg.Add(-1)과 같다고 보시면 됩니다.   

그렇기 때문에 처음에 WaitGroup을 만들자마자 URL 개수만큼 카운터를 증가시키고 각각의 고루틴에서는 작업이 완료될 때까지 카운터를 감소시킵니다. 모든 고루틴이 끝나면 
카운터가 0이 되는데, 이 상태가 되기전까지는 wg.Wait() 부분에서 멈춰있게 됩니다.

미리 고루틴이 몇 개 생길지 알기 때문에 이렇게 작성이 가능합니다만, 고루틴이 몇개 생길지 알기 어렵거나 따로 수를 세어야 알 수 있는 경우에는 고루틴 띄워보내기 전에 
wg.Add(1)을 수행하여 하나씩 카운터를 증가시킬 수 있습니다.

```go
package main

import "log"

func main() {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			if _, err := download(url); err != nil {
				log.Fatal(err)
			}
		}(url)
	}
	wg.Wait()
}
```

## 공유 메모리와 최솟값 찾기 

공유메모리를 이용하여 매우 병렬화가 잘되는 문제를 풀어보기 

```go
package main

import (
	"fmt"
	"sync"
)

func FindMinimumValue() {
	fmt.Println(parallelMin(
		[]int{
			83, 46, 49, 23, 97, 12, 11, 46, 49, 27, 5,
		}, 2))
}

func parallelMin(a []int, n int) int {
	if len(a) < n {
		return min(a)
	}

	mins := make([]int, n)
	size := (len(a) + n - 1) / n
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			begin, end := i*size, (i+1)*size
			if end > len(a) {
				end = len(a)
			}
			mins[i] = min(a[begin:end])
		}(i)
	}
	wg.Wait()
	return min(mins)
}

func min(a []int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]

	for _, e := range a[1:] {
		if min > e {
			min = e
		}
	}
	return min
}

```

## Channel 

채널도 First Class Citizen 이다.  
채널은 양방향 채널이 있고, 단방향 채널이 있습니다.  
양방향 채널은 자연스럽게 단방향 채널로 변환해서 쓸 수 있습니다. 

```go
package main

func main()  {
	c1 := make(chan int)
	var c2 chan int= c1 // c1 과 c2는 동일한 채널이다. 
	var c3 <-chan int= c1  // 자료를 뺄수만 잇는 채널의 자료형 
	var c4 chan<- int = c1   // 자료를 넣을수만 있는 채널의 자료형
}
```

## 일대일 단방향 채널 소통 

- channel이 호출되는 갯수를 알고 있는 경우 

```go
package main

import "fmt"

func Example_simpleChannel() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
	}()

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
```

- channel이 호출되는 갯수를 모르는 경우

채널 하나를 만들어서 넘겨주고 넘겨받는 것이 깔끔해보이지 않기 때문에 주로 함수가 채널을 반환하게 만드는 패턴을 쓰게 됩니다. 

```go
package main

import "fmt"

func Example_simpleChannelWithDynamicCount() {
	c := func() <-chan int {
		c := make(chan int)
		go func() {
			defer close(c) // 보내는 쪽에서 close(c)로 채널을 마지막에 닫아주었음.
			c <- 1
			c <- 2
			c <- 3
		}()
		return c
	}()

	for num := range c {
		fmt.Println(num)
	}
}

``` 

- 생성기 패턴 ( 채널 활용 )
  - 생성하는 쪽에서는 상태 저장 방법을 복잡하게 고민할 필요가 없다. 
  - 받는 쪽에서는 for의 range를 이용할 수 있다. 
  - 채널 버퍼를 이용하면 멀티 코어를 활용하거나 입출력 성능 상의 장점을 이용할 수 있다. 

```go
package main

import "fmt"

func FibonacciWithChannel(max int) <-chan int {
	c := make(chan int)

	go func() {
		defer close(c)
		a, b := 0, 1
		for a <= max {
			c <- a
			a, b = b, a+b
		}
	}()
	return c
}

func ExampleFibonacci() {
	
	fmt.Println("FibonacciWithChannel")

	for fib := range FibonacciWithChannel(15) {
		fmt.Print(fib, ",")
	}

	fmt.Println(" ")
}

```

- 버퍼 있는 채널 

버퍼가 있는 채널은 아래와 같이 구성할 수 있다.  
버퍼의 장점은 보내는 쪽과 받는 쪽의 코드가 균일한 속도로 수행되지 않는 경우입니다. 사실 입출력 역시 한꺼번에 이루어지는 경우가 많지만, 
입출력은 균일한 속도로 이루어진다고 해도, 패턴에 맞는 결과가 몰려 있을 수도 있고 듬성듬성 나타날 수도 있습니다. 이 때 버퍼를 만들어주면, 
이두 고루틴 간에 어느 정도 격차가 생겨도 계속 동작할 수 있기 때문에 성능 향상이 일어날 수 있습니다.  

**동시성은 강력하지만 복잡할 수 있으므로 알려진 패턴을 따르시는 것이 좋습니다. 버퍼 없는 채널로 동작하는 코드를 만들고 필요에 따라 성능 향상을 위하여 버퍼 값을 조절하는 것이 좋겠습니다.** 

```go
package main

func main()  {
	c := make(chan int, 10)
    
}
```

- 닫힌 채널 

채널이 닫혀잇는 경우에는 ?

```go
package main

import "fmt"

func main() {

	c := make(chan int, 10)

	close(c)
	
	val, ok := <-c

	if ok {
		fmt.Println("Channel 이 열려있습니다. ")
	}
}
```

하나의 변수로 받거나 두 변수로 받거나 상관 없이 val 에는 기본 값이 들어옵니다, 기본 값을 0, 빈 문자열, nil, 모든 필드가 기본값으로 되어 있는 구조체 등이되며, ok에는 false 값으 넘어옵니다.   
**닫은 채널은 또 닫으면 패닉이 발생합니다**  
이를 활용해서 다양한 에러 처리 방법을 강구해볼 수 있습니다. 

## 동시성 패턴 

### 파이프라인 패턴

파이프라인은 한 단계의 출력이 다음 단계의 입력으로 이어지는 구조 입니다. 
컴퓨터 시스템에서는 특히 서로 다른 종류의 하드웨어들이 어떤 일을 해야할 때 파이프라인이 큰 효과가 있습니다. 소프트웨어에서는 들어오는 데이터와 나가는 데이터에 
집중하여 문제를 풀고자 할 때 장점이 있고, 버퍼를 활용하면 경우에 따라 성능상의 장점도 얻을 수 있습니다. 

- 기본적인 파이프라인 패턴 

```go
package main

import "fmt"

func PlusOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num + 1
		}
	}()
	return out
}

func ExamplePlusOne() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()

	for num := range PlusOne(PlusOne(c)) {
		fmt.Println(num)
	}
}
```

- chain으로 이어진 파이프라인을 구성해야할 경우 

```go
package main

import "fmt"

func PlusOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num + 1
		}
	}()
	return out
}

type InitPipe func(<-chan int) <-chan int

func Chain(ps ...InitPipe) InitPipe {
	return func(in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(c)
		}
		return c
	}
}

func ExampleWithChain() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()

	for num := range Chain(PlusOne, PlusOne)(c) {
		fmt.Println(num)
	}
}
```

### 채널 공유로 팬아웃 하기 

- Fan-Out : 논리회로에서 주로 쓰이는 용어로 게이트 하나의 출력이 게이트 여러 입력으로 들어가는 경우를 팬아웃이라고 합니다.

```go
package main

import (
	"fmt"
	"time"
)

func FanoutMain() {
	c := make(chan int)

	for i := 0; i < 3; i++ {
		go func(i int) {
			for n := range c {
				time.Sleep(1)
				fmt.Println(i, n)
			}
		}(i)  
		// 3개의 고루틴을 호출 할 때 i를 따로 넘겨주었습니다. 그리고 각각의 고루팀은 i 를 받아서 사용했습니다. 
		// 꼭 이렇게 해주어야 하는데, 그렇게 하지않으면 고루틴이 수행되는 시점에 i값을 가져다가 사용하게 됩니다.
		// 메인 고루틴에서 i값을 계속 증가시키기 때문에 증가된 값을 가져다가 쓰게 되어서 잘못된 값이 사용됩니다. 
	}
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}
```

채널을 닫지 않아도 프로그램이 종료되지만 만일 여기서 프로그램이 종료되지 않는 경우에는 숫자들을 기다리는 3개의 고루틴이 종료되지 않아서 계속 메모리에 남아 있게 됩니다. 
오랫 동안 동작하는 서버 프로그램에서 이런 실수를 하게 되면 시간이 흐를 수록 고루틴의 수가 점점 늘어나게 되어서 메모리 누수가 발생하게 됩니다.    

그리고 팬아웃에서 채널은 닫는 것은 방송(broadcast) 효과가 있습니다. 만일 채널을 닫는 대신 미리 약속한 특수한 값을 채널로 전달하여 종료되었다는 것을 알린다면 
채널에서 자료를 받는 고루틴 중에서 하나만 이것을 전달 받게 됩니다. 그러면 받은 고루틴의 수만큼 이 특수한 값을 보내주어야 하는데, 채널을 닫아버리면 몇 개의 고루틴이 
이 채널에서 값이 오기를 기다리는지에 상관 없이 모두 이것을 알게 됩니다. 따라서 채널을 닫는 것은 신호를 모두에게 전달하기 위한 매우 강력하고 깔끔한 방법입니다. 

### 팬인하기

- Fan-In : 논리회로에서 주로 쓰이는 용어로 하나의 게이트에 여러개의 입력선이 들어가는 경우를 팬인이라고 합니다. 

```go
package main

import (
	"fmt"
	"sync"
)

func FanIn(ins ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ins))
	for _, in := range ins {
		go func(in <-chan int) {
			defer wg.Done()
			for num := range in {
				out <- num
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func ExampleFanIn() {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	c3 := make(chan int, 10)

	for i := 0; i < 10; i++ {
		c1 <- i
	}
	for i := 0; i < 10; i++ {
		c2 <- i
	}
	for i := 0; i < 10; i++ {
		c3 <- i
	}

	channels := FanIn(c1, c2, c3)

	// todo - 고루틴을 사용하지 않으면 해당 문장에서 멈춰버림!!
	go func() {
		for num := range channels {
			fmt.Print(num)
		}
	}()
}
```

### 분산처리 

- 팬아웃 해서 파이프라인을 통과시키고 다시 팬인 시키면 분산 처리가 됩니다. 

고루틴의 갯수가 많은 것은 크게 걱정할 필요가 없습니다. Go에서는 고루틴 마다 스레드를 모두 할당하지 않으며, 동시에 수행될 필요가 없는 고루틴들은 모두 하나의 스레드에서 순차적으로 수행되며 이것이 컴파일 시간이 
예측가능한 경우가 많으므로 스레드를 이렇게 많이 만드는 경우 생길 수 있는 비용이 발생하지 않습니다. 

```go
package main

import (
	"fmt"
	"sync"
)

func Distribute(p InitPipe, n int) InitPipe {
	return func(ints <-chan int) <-chan int {
		cs := make([]<-chan int, n)
		for i := 0; i < n; i++ {
			cs[i] = p(ints)
		}
		return FanIn(cs...)
	}
}

func ExampleDistribute() {
	fmt.Println("------------ Start ExampleDistribute ------------")
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()

	var wg sync.WaitGroup

	out := Chain(PlusOne, Distribute(Chain(PlusOne, PlusOne, PlusOne), 10), PlusOne)(c)

	wg.Add(1)

	go func() {
		for num := range out {
			fmt.Println(num)
		}

		wg.Done()
	}()

	wg.Wait()

	fmt.Println("------------ End ExampleDistribute ------------")
}
```

### select 

select 를 이용하면 동시에 여러 채널과 통신할 수 있습니다. select의 형태는 switch과 비슷하지만 아래의 특징이 있습니다.

- 모든 case가 계산된다. 거기에 함수 호출등이 있으면 select 수행할 때 모두 호출된다. 
- 각 case는 채널에 입출력하는 현태가 되며 막히지 않고 입출력이 가능한 case가 있으면 그 중에 하나가 선택되어 입출력이 수행되고 해당 case의 코드만 수행된다. 
- default가 있으면 모든 case에 입출력이 불가능할 때 코드가 수행된다.

```go
package main

import "fmt"

func main() {
  select {
    case n := <-c1:
        fmt.Println(n, "is from c1")
    case n := <-c1:
        fmt.Println(n, "is from c1")
    case c3 <- f():
        fmt.Println(n, "is from c1")
    default :
	    fmt.Println("No Channel is ready")
  }
}

```

### select - 팬인하기 

이 코드의 특징은 닫힌 채널을 nil 로 바꾸어주었다는 데 있습니다. nil 채널에는 보내기 및 받기가 모두 막히게 됩니다. 그렇기 때문에 채널이 닫혔다는 것이 
발견되면 이것을 영원히 채널로 바꿔준 건입니다. 물론 채널 자체를 바꾼 것이 아니라 채널 별수를 nil 로 바꾼 것이기 때문에 혹시나 이 예제에는 보이지 않는 
다른 고루틴이 닫힌 채널에서 자료를 받아가고 있었다고 해도 그쪽에는 아무런 영향을 주지 않습니다. 

```go
package main

import "fmt"

func FanInSelect(in1, in2, in3 <-chan int) <-chan int {

	out := make(chan int)
	openCnt := 3

	closeChan := func(c *<-chan int) bool {
		*c = nil
		openCnt--
		return openCnt == 0
	}

	go func() {
		defer close(out)
		for {
			select {
			case n, ok := <-in1:
				if ok {
					out <- n
				} else if closeChan(&in1) {
					return
				}
			case n, ok := <-in2:
				if ok {
					out <- n
				} else if closeChan(&in2) {
					return
				}
			case n, ok := <-in3:
				if ok {
					out <- n
				} else if closeChan(&in3) {
					return
				}
			}
		}
	}()
	return out
}

func Example_SelectFanIn() {
	c1, c2, c3 := make(chan int), make(chan int), make(chan int)
	sendInts := func(c chan<- int, begin, end int) {
		defer close(c)
		for i := begin; i < end; i++ {
			c <- i
		}
	}

	go sendInts(c1, 11, 14)
	go sendInts(c2, 21, 23)
	go sendInts(c3, 31, 35)

	for n := range FanInSelect(c1, c2, c3) {
		fmt.Println(n, ",")
	}
}

```

#### select - 채널을 기다리지 않고 받기

채널에 값이 있으면 받고 없으면 그냥 스킵하는 흐름을 구성하려면 

```go
package main

import "fmt"

func main() {
  select {
  case n := <-c:
    fmt.Println(n)
  default:
    fmt.Println("Data is not ready. Skipping...")
  }
}
```

### select - 시간 제한 

채널과 통신을 기다리되 일정 시간 동안만 기다리겠다면 time.After 함수를 이용할 수 있습니다.  
아래와 같은 코드가 작성되면 recv 와 send에 빈번하게 자료가 반복적으로 오고 가더라도 5초 동안만 처리하게 됩니다. 

```go
package main

import (
  "fmt"
  "time"
)

func Main_TimeLimit(recv <-chan int, send chan<- int ) {
  timeout := time.After(5 * time.Second)
  for {
    select {
    case n := <-recv:
        fmt.Println(n)
    case send <- 1:
		fmt.Println("sent 1")
    case  <- timeout:
        fmt.Println("communication wasn't finished in 5 sec")
		return
    }
  }
}
```

### 파이프라인 중단하기

이전에는 파이프라인을 구성할 때, 받기만 할 뿐 그만 보내달라고 요청할 수 없었습니다.  
모두 자료를 소진시키지 않으면 해제되지 않는 고루틴들이 메모리에 남아 있으면 메모리 누수가 발생합니다. 
모든 자료를 소진시킨다고 해도 좋지 않는 경우가 많은데 보내는 고루틴이 많은 네트워크 트래픽을 유발시키거나 배터리를 소모한다면 
계속해서 데이터를 받아오면 그만큼 더 많은 네트워크 트래픽과 배터리 소모가 발생할 것이기 때문입니다.   

- 이럴 때 유용한 패턴 중 하나는 done 채널을 하나 더 두는 것입니다. 보내는 고루틴에서 이 채널로부터 신호가 감지되면 보내는 것을 중단하고 채널을 닫으면 됩니다. 신호는 close(done)으로 주면 됩니다. 

```go
package main

import (
  "fmt"
  "runtime"
  "time"
)

func PlusOneWithChannel(done <-chan struct{}, in <-chan int) <-chan int {
  out := make(chan int) // 양방향으로 채널이 사용될 수 있음
  go func() {
    defer close(out)
    for num := range in {
      select {
      case out <- num + 1:
      case <-done:
        return
      }
    }
  }()
  return out
}

func ReturnStopPipelineWithChannel() {
  c := make(chan int)
  go func() {
    defer close(c)
    for i := 3; i < 103; i += 10 {
      c <- i
    }
  }()

  done := make(chan struct{})
  nums := PlusOneWithChannel(done, PlusOneWithChannel(done, PlusOneWithChannel(done, PlusOneWithChannel(done, PlusOneWithChannel(done, c)))))

  for num := range nums {
    fmt.Println(num)
    if num == 18 {
      break
    }
  }

  close(done)
  time.Sleep(100 * time.Millisecond)

  fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
  for _ = range nums {
    // Consume All nums
  }
  time.Sleep(100 * time.Millisecond)
  fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
}
```

PlusOneWithChannel의 반복문에서 select를 이용하여 done 채널도 함께 관착하는 것을 알 수 있습니다. select에서 닫힌 채널로부터 
부한정 기본값(빈 값)을 받아오는 것이 다소 불편해보였는데, 여기서는 좋은 특성이 됩니다. 만약에 이런 특성이 없엇다면 done 채널에 값에 기다르는 회수에 
맞게 done 채널에 값을 보내줘야하는 데 이것은 쉬운일이 아닙니다. close(done) 한 번으로 이 채널로부터 값을 기다리고 있는 모든 로그틴에 일이 끝났다고 
방송을 하는 것입니다. 

### 컨텍스트(context.Context) 활용하기 

done 채널을 따로 운영해도 좋습니다만 이뿐만 아니라 더 복잡한 상황이 발생하기 때문에 context* 패턴을 이용하면 좋습니다. 복잡한 상황이란 여러 고루틴에 
종료 신호 이외에도 다른 공유되어야 하는 정보가 있다는 것입니다. 대표적으로 사용자 인증 정보나 요청 마감등이 있습니다.   

context.Context는 계층 구조로 되어 있습니다. context.Background()가 가장 상위에 있습니다. 이것은 프로그램이 끝날 때 까지 절대로 취소되지 않고 
계속 살아 있습니다. 여기에 하위 구조를 계속해서 트리 구조로 붙일 수 있는데 상위 구조가 취소되면 그 하위에 있는 모든 컨텍스트도 취소됩니다.   

context.WithCancel로 취상위 컨텍스트인 context.Background() 밑에 취소 기능을 갖춘 컨텍스트를 붙였습니다. ctx와 cancel 두 변수로 받았는데 ctx
에는 새로 생성된 컨텍스트가 들어가고, cancel 은 이 컨텍스트를 취소하는데 호출할 수 있는 함수가 들어 갑니다.  

---

WithDeadline, WithTimeout 을 이용하여 만든 ctx를 이용하여 호출하면 시간이 지나면 취소되게 만들 수 있습니다. 이것을 활용하지 않고 구현한다면 상당히 번거
롭습니다. WithValue를 이용하면 인증 토큰 같이 요청 범위 내에 있는 값들을 보낼 수 있어서 편리합니다. 

> [golang - context](http://godoc.org/golang.org/x/net/context)

### 요청과 응답 짝짓기 

요청을 한 채널에 보내고 응답을 다른 채널로 받는 방식으로 파이프라인을 동작시킬 수 있을 것입니다. PlusOne의 경우에는 넘겨준 채널에 1을 넣으면 반환받는 채널에서 2를
받을 수 있으니까요. 그런데 고민이 될 수 있는 점이 있습니다. 응답을 받았을 때 이것이 어느 요처엥 의한 응답인지 알아야 하는 경우가 있습니다. 특히 분산 처리되면 어느것이
먼저 나올지 알 수 없데 되니까요. 알 필요가 없는 경우도 많지만 알아야 하는 경우도 있습니다.

```go
package main

import (
	"fmt"
	"sync"
)

type Request struct {
	Num  int
	Resp chan Response
}

type Response struct {
	Num      int
	WorkerID int
}

func PlusOneService(reqs <-chan Request, workerId int) {
	for req := range reqs {
		go func(req Request) {
			defer close(req.Resp)
			req.Resp <- Response{req.Num + 1, workerId}
		}(req)
	}
}

func MappingRequestAndResponse() {
	fmt.Println("Start -------------- MappingRequestAndResponse --------------")
	
	reqs := make(chan Request)
	defer close(reqs)
	for i := 0; i < 3; i++ {
		go PlusOneService(reqs, i)
	}
	var wg sync.WaitGroup
	for i := 3; i < 53; i += 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			resps := make(chan Response)
			reqs <- Request{i, resps}
			fmt.Println(i, "=>", <-resps)
		}(i)
	}
	wg.Wait()
}
```
### 동적으로 고루틴 이어붙이기

```go
package main

import (
	"context"
	"fmt"
)

type IntPipe func(context.Context, <-chan int) <-chan int

func Range(ctx context.Context, start, step int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; ; i += step {
			select {
			case out <- i:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func FilterMultiple(n int) IntPipe {
	return func(ctx context.Context, in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for x := range in {
				if x%n == 0 {
					continue
				}
				select {
				case out <- x:
				case <-ctx.Done():
					return
				}
			}
		}()
		return out
	}
}

func Primes(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		c := Range(ctx, 2, 1)
		for {
			select {
			case i := <-c:
				c = FilterMultiple(i)(ctx, c)

				select {
				case out <- i:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func PrintPrimes(max int) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for prime := range Primes(ctx) {
		if prime > max {
			break
		}
		fmt.Println(prime, " ")
	}
	fmt.Println()
}
```

# 주의해야 할 사항 

- 보내는 쪽에서 채널을 닫는 것이 매우 중요합니다.

```go
package main

import "fmt"

func main() {
  c := make(chan int)
  done := make(chan bool)

  go func() {
    for i := 0; i < 10; i++ {
      c <- i
    }
    done <- true
  }()
  go func() {
    for {
      fmt.Println( <- c)
    }
  }()
  <- done
}
```

- 첫번째 케이스 

위의 코드는 여러가지 문제점이 있습니다. 일단 바로 눈여겨볼 것은 두 번째 고루틴이 끝나지 않는다는 것입니다. 위의 코드가 반복적으로 수행될 경우에 고루틴의 수는 점점 늘어날 것입니다.

- 두번째 케이스 

첫번째 고루틴인 생산자에서는 생산이 끝난 다음에 done에 true 값을 넣어줍니다. 그리고 메인 고루틴은 <- done 으로 이것이 끝날 때 까지 기다립니다.
여기서 done 은 생산이 끝난 뒤에 값이 들어가므로 소비가 끝나기 전에 메인이 고루틴이 끝나벌릴 가능성이 있습니다. 위에서 아무리 fmt.Println(<-c)가 
한줄이라고 해도 <-c 로 받아와서 출력을 할 때 문맥 전환이 일어날 수 있는 것이지요. 혹은 이 다음 줄에 다른 일을 더 해야 할 수도 있을 것입니다. 
따라서 운이 나쁘면 이 코드는 마지막 하나가 출력이 되지 않을 수 있습니다. 이렇게 코드를 만들면 나중에 버그가 발생했을 때 이유를 찾기 어렵습니다. 

## 정리 

- 자료를 보내는 채널은 보내는 쪽에서 닫는다. 
- 보내는 쪽에서 반복문 등을 활용해서 보내다가 중간에 return을 할 수 있으므로 닫을 때는 defer를 이용하는 것이 좋다. 그렇지 않으면 중간에 return 했을 때 채널을 닫지 않고 종료할 수 있다. 
- 받는 쪽이 끝날 때 까지 기다리는 것이 모든 자료의 처리가 끝나는 시점까지 지다리는 방법으로 더 안정적이다. 위의 예제에서는 생산자 쪽이 아닌 소비자 쪽에서 done <- true를 했어야 했다. 물론 위의 예제에서는 소비자 쪽에서 언제 끝났는지 알 수 없다. 그것을 생산자에서 채널을 다는 것으로 신호를 줬어야 했다. 
- 특별한 이유가 없다면 받는 쪽에서는 range를 이용하는 것이 좋다. 생산자가 채널을 다는 경우에 반복문을 빠져 나오게 되기 때문에 편리하다. 
- 루틴이 끝났음을 알리고 다른 쪽에서 기다리는 것은 sync.WaitGroup 을 이용하는 것이 나은 경우가 더 많다. 위의 예제와 같이 간단한 경우가 아니라면 두 고루틴 모두가 끝날 때까지 기다렸으면 문제가 없을 것이다. 
- 끝났음을 알리는 done 채널은 자료를 보내는 쪽에서 결정할 사항이 아니다. 자료를 보내는 쪽에서 채널을 닫아서 자료가 끝났음을 알리는 것이 더 낫다. 그러면 done 채널은 받는 쪽에서 보내는 쪽으로 자료 전송이 끝났거나 끝나지 않았으면 더 이상 자료를 보내지 말아달라는 cancel 요청으로 보는 것이 낫다. 
- done 채널에 자료를 보내어 신호를 주는 많은 예제가 있는데, close(done) 으로 채널을 다는 것이 더나은 방법인 경우가 많다. 

