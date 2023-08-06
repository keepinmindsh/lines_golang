package main

import (
	car "design_pattern/oop/app/car"
	maps "design_pattern/oop/app/maps"
	"design_pattern/oop/domain"
)

func main() {
	newMaps := maps.NewMaps(10, 10)

	var vector domain.Vector

	steering := car.NewSteering(vector)

	myCar := car.NewCar(car.NewTire(car.NEXEN, &newMaps, steering), steering)

	myCar.Start()

	myCar.LeftDiagonal()
	myCar.MoveBackward()
	myCar.MoveForward()

	myCar.Stop()
}
