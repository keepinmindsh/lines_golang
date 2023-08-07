package main

import (
	"design_pattern/oop/app/car/service/car"
	"design_pattern/oop/app/car/service/moving"
	"design_pattern/oop/app/car/service/steering"
	"design_pattern/oop/app/car/service/tire"
	maps "design_pattern/oop/app/maps"
	"design_pattern/oop/domain"
)

func main() {
	newMaps := maps.NewMaps(10, 10)

	var vector domain.Vector

	steering := steering.NewSteering(vector)
	moving := moving.NewMoving(domain.TICO, &newMaps)

	myCar := car.NewCar(tire.NewTire(tire.NEXEN, steering, moving), steering, moving)

	myCar.Start()

	myCar.LeftDiagonal()
	myCar.MoveForward()

	myCar.Stop()
}
