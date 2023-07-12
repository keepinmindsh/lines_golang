package car

import "designpattern_v1/v1/pattern/srp/domain"

type volvo struct {
}

func NewCar() domain.Car {
	return volvo{}
}

func (v volvo) MoveForward() {
}

func (v volvo) MoveBackward() {
}

func (v volvo) Stop() {
}

func (v volvo) StartDriving() {
}
