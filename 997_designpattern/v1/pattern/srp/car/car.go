package car

import "designpattern_v1/v1/pattern/srp/domain"

type Car struct {
	domain.Tire
	domain.Steering
}

func NewCar(tire domain.Tire, steering domain.Steering) Car {
	return Car{tire, steering}
}
