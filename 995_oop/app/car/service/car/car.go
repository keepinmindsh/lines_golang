package car

import (
	"design_pattern/oop/domain"
)

type RealCar struct {
	domain.Tire
	domain.Steering
	domain.Moving
}

func NewCar(tire domain.Tire, steering domain.Steering, moving domain.Moving) domain.Car {
	return RealCar{
		Tire:     tire,
		Steering: steering,
		Moving:   moving,
	}
}
