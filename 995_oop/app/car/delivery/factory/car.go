package factory

import (
	"design_pattern/oop/domain"
)

type RealCar struct {
	domain.Tire
	domain.Steering
}

func NewCar(tire domain.Tire, steering domain.Steering) domain.Car {
	return RealCar{
		Tire:     tire,
		Steering: steering,
	}
}
