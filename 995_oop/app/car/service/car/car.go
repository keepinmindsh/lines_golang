package car

import (
	carDomain "design_pattern/oop/domain/car"
	steeringDomain "design_pattern/oop/domain/steering"
	tireDomain "design_pattern/oop/domain/tire"
)

type RealCar struct {
	tireDomain.Tire
	steeringDomain.Steering
}

func NewCar(tire tireDomain.Tire, steering steeringDomain.Steering) carDomain.Car {
	return RealCar{
		Tire:     tire,
		Steering: steering,
	}
}
