package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
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
	fmt.Println("PlusOneService Start : " + time.Now().Format("YYYY MM DD HH24:MI:SS"))
	for req := range reqs {
		go func(req Request) {
			defer close(req.Resp)
			req.Resp <- Response{req.Num + 1, workerId}
			fmt.Println("PlusOneService Start - go func : " + time.Now().Format("YYYY MM DD HH24:MI:SS") + " Value : " + strconv.Itoa(req.Num))
		}(req)
	}
}

// todo 이 코드는 정말 이해가 되지 않습니다 chan은 정의되는 순간부터 동일한 주소로 인식될까요?
// 재생각엔 request가 수신 채널로 받아졌을 때, 해당 이벤트의 트리거로 송신 채널 req가 열렸다고 보는게 제일 맞음요.
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
			fmt.Println(&resps)
			reqs <- Request{i, resps}
			fmt.Println(i, "=>", <-resps)
		}(i)
	}
	wg.Wait()
}
