package sample

import (
	"fmt"
	"time"
)

func TimeTimerForSeconds() {
	timer1 := time.NewTimer(3 * time.Second)
	done := make(chan bool)
	go func() {
		<-timer1.C
		fmt.Println("Timer 1 fired")
		done <- true
	}()
	<-done
}

func TimeSimpleTickerByOneSecond() {
	fmt.Println("Go Tickers Tutorial")
	// this creates a new ticker which will
	// `tick` every 1 second.
	ticker := time.NewTicker(1 * time.Second)

	// for every `tick` that our `ticker`
	// emits, we print `tock`
	for t := range ticker.C {
		fmt.Println("Invoked at ", t)
	}
}
