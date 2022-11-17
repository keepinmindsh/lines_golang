package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"time"
)

func secondTask() {

	fmt.Println("Task:", time.Now())
}

func hourTask() {
	fmt.Println("Hour Task:", time.Now())
}

func main() {

	fmt.Println("go Scheduler example")

	gocron.Every(1).Second().Do(secondTask)

	// Begin job immediately upon start
	gocron.Every(1).Hour().From(gocron.NextTick()).Do(hourTask)

	//gocron.Every(1).Hour().Do(hourTask)

	// Start all the pending jobs
	<-gocron.Start()

	// also, you can create a new scheduler
	// to run two schedulers concurrently
	s := gocron.NewScheduler()
	<-s.Start()

}
