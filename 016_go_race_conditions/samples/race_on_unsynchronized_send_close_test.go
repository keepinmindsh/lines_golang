package samples

import "testing"

func Test_UnSynchronizedSendClose(t *testing.T) {

}

func WrongCase() {
	c := make(chan struct{})

	go func() { c <- struct{}{} }()
	close(c)
}

// RightCase
// 고의 메모리 모델에 따르면, 채널로 부터 완료를 수신받아 처리되기 전에 하나의 채널의 전송이 발생한다.
// 전송과 닫기를 동기하하기 위해서는 수신 프로세스를 사용해서 채널을 닫기전에 전송이 완료되었음을 반드시 보장해야한다.
func RightCase() {
	c := make(chan struct{})

	go func() { c <- struct{}{} }()
	<-c
	close(c)
}
