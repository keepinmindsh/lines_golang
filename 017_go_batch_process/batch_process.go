package main

import (
	"flag"
	"fmt"
	batch "github.com/Deeptiman/go-batch"
	log "github.com/sirupsen/logrus"
)

// Resources Structure
type Resources struct {
	id   int
	name string
	flag bool
}

func main() {
	var rFlag, mFlag int
	flag.IntVar(&rFlag, "r", 1000, "No of resources")
	flag.IntVar(&mFlag, "m", 100, "Maximum items")
	flag.Parse()

	logs := log.New()

	logs.Infoln("Batch Processing Example !")

	b := batch.NewBatch(batch.WithMaxItems(uint64(mFlag)))

	b.StartBatchProcessing()

	// go func() {

	// 	select {
	// 	case <-time.After(time.Duration(2) * time.Second):
	// 		fmt.Println("Run Batch processing again!")
	// 		b.StartBatchProcessing()
	// 	}

	// }()

	go func() {

		// Infinite loop to listen to the Consumer Client Supply Channel that releases
		// the []BatchItems for each iteration.
		for {
			for bt := range b.Consumer.Supply.ClientSupplyCh {
				logs.WithFields(log.Fields{"Batch": bt}).Warn("Client")
			}
		}
	}()

	for i := 1; i <= rFlag; i++ {
		b.Item <- &Resources{
			id:   i,
			name: fmt.Sprintf("%s%d", "R-", i),
			flag: false,
		}
	}
	b.Close()

}

func basicSample() {
	b := batch.NewBatch(batch.WithMaxItems(100))
	go b.StartBatchProcessing()

	for i := 1; i <= 1000; i++ {
		b.Item <- &Resources{
			id:   i,
			name: fmt.Sprintf("%s%d", "R-", i),
			flag: false,
		}
	}
	b.Close()
}
