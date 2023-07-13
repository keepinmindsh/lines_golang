package main

import (
	"designpattern_v1/v1/pattern/srp/car"
	"designpattern_v1/v1/pattern/srp/maps"
	"fmt"
)

func main() {
	newMaps := maps.NewMaps(5, 5)

	newCar := car.NewCar(&newMaps)

	newCar.StartDriving()
	newCar.MoveForward()
	newCar.MoveForward()
	newCar.MoveForward()
	newCar.MoveForward()
	newCar.MoveForward()
	newCar.MoveBackward()
	newCar.Stop()

	fmt.Println(newMaps)
}
