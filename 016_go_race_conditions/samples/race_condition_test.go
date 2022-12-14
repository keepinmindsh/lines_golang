package samples

import (
	"fmt"
	"testing"
)

func Test_RaceCondition(t *testing.T) {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a"
		c <- true
	}()
	m["2"] = "b"
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}
