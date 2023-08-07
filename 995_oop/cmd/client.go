package main

import (
	"design_pattern/oop/app/car/delivery/factory"
	maps "design_pattern/oop/app/maps"
	"design_pattern/oop/domain"
)

func main() {
	newMaps := maps.NewMaps(10, 10)

	var vector domain.Vector

	steering := factory.NewSteering(vector)

	myCar := factory.NewCar(factory.NewTire(factory.NEXEN, &newMaps, steering), steering)

	myCar.Start()

	myCar.LeftDiagonal()
	myCar.MoveBackward()
	myCar.MoveForward()

	myCar.Stop()
}
