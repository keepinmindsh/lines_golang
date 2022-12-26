package sample

import "fmt"

func CountDown() {

	for count := 10; count >= 0; count-- {
		if count == 10 {
			fmt.Printf("Ignition Sequence Start!")
		} else {
			fmt.Printf(`%v ....`, count)
		}
	}

}
