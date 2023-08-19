package tire

import (
	"design_pattern/oop/app/car/service/moving"
	steeringDomain "design_pattern/oop/domain/steering"
	tireDomain "design_pattern/oop/domain/tire"
	"design_pattern/oop/internal/code"
	"fmt"
)

type nexenTire struct {
	moving     *moving.Moving
	steering   steeringDomain.Steering
	TireStatus tireDomain.Status
}

func NewNexenTire(moving *moving.Moving, steering *steeringDomain.Steering) tireDomain.Tire {
	return &nexenTire{
		moving:   moving,
		steering: *steering,
	}
}

func (v *nexenTire) Forward() {
	v.TireStatus = tireDomain.FOWARD

	currentVector := v.steering.CurrentVector()

	switch currentVector {
	case code.RightDiagonal:
		v.moving.RightForwardDiagonal.Move()
	case code.LeftDiagonal:
		v.moving.LeftForwardDiagonal.Move()
	case code.Straight:
		v.moving.Forward.Move()
	}
}

func (v *nexenTire) Backward() {
	v.TireStatus = tireDomain.BACKWARD

	v.moving.Backward.Move()
}

func (v *nexenTire) Stop() {
	v.TireStatus = tireDomain.STOP

	fmt.Println("정지")
}

func (v *nexenTire) Start() {
	v.TireStatus = tireDomain.START

	fmt.Println("부릉")
}
