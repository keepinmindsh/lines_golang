package car

import "oop/domain"

type Car struct {
	domain.Tire
	domain.Steering
}

func NewCar(tire domain.Tire, steering domain.Steering) domain.Car {
	return Car{tire, steering}
}
