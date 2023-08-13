package car

import (
	carDomain "design_pattern/oop/domain/car"
	movingDomain "design_pattern/oop/domain/moving"
	steeringDomain "design_pattern/oop/domain/steering"
	tireDomain "design_pattern/oop/domain/tire"
)

type RealCar struct {
	tireDomain.Tire
	steeringDomain.Steering
	movingDomain.Moving
}

func NewCar(tire tireDomain.Tire, steering steeringDomain.Steering, moving movingDomain.Moving) carDomain.Car {
	return RealCar{
		Tire:     tire,
		Steering: steering,
		Moving:   moving,
	}
}
