package main

import (
	"design_pattern/oop/app/car/service/car"
	"design_pattern/oop/app/car/service/moving"
	steeringSvc "design_pattern/oop/app/car/service/steering"
	"design_pattern/oop/app/car/service/tire"
	maps "design_pattern/oop/app/maps"
	"design_pattern/oop/app/maps/validate"
	"design_pattern/oop/domain"
	"design_pattern/oop/domain/steering"
	"design_pattern/oop/internal/logger"
)

func main() {
	logger.InitLogger()

	newMaps := maps.NewMaps(10, 10)

	var vector steering.Vector

	steering := steeringSvc.NewSteering(vector)
	moving := moving.NewMoving(domain.TICO, &newMaps)

	validate.NewValidater(&newMaps)

	myCar := car.NewCar(
		tire.NewTire(tire.NEXEN, moving, &steering),
		steering,
	)

	myCar.Start()

	myCar.LeftDiagonal()
	myCar.Forward()
	myCar.Forward()
	myCar.RightDiagonal()
	myCar.Forward()
	myCar.Forward()

	myCar.Stop()
}
