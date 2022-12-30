package sample

import (
	"fmt"
	"time"
)

func StoppingTicker() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
