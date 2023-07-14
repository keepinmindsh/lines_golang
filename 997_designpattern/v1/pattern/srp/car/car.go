package car

import "designpattern_v1/v1/pattern/srp/domain"

type Car struct {
	domain.Tire
}

func NewCar(tire domain.Tire) Car {
	return Car{tire}
}
