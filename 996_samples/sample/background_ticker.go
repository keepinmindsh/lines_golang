package sample

import (
	"fmt"
	"time"
)

func bgTask() {
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		fmt.Println("Tock")
	}
}

func BackgroundTimer() {
	fmt.Println("Go Tickers Tutorial")

	go bgTask()

	fmt.Println("The rest of my application can continue")

	select {}
}
