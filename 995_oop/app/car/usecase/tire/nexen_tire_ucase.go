package tire

import (
	steeringMoving "design_pattern/oop/domain/moving"
	"design_pattern/oop/domain/steering"
	tireDomain "design_pattern/oop/domain/tire"
	"fmt"
)

type nexenTire struct {
	steering   steering.Steering
	moving     steeringMoving.Moving
	TireStatus tireDomain.Status
}

func NewNexenTire(steering steering.Steering, moving steeringMoving.Moving) tireDomain.Tire {
	return &nexenTire{
		steering: steering,
		moving:   moving,
	}
}

func (v *nexenTire) Forward() {
	v.TireStatus = tireDomain.FOWARD

	v.moving.Move(steeringMoving.MoveOrder{
		Vector:     v.steering.CurrentVector(),
		TireStatus: tireDomain.FOWARD,
	})
}

func (v *nexenTire) Backward() {
	v.TireStatus = tireDomain.BACKWARD

	v.moving.Move(steeringMoving.MoveOrder{
		Vector:     v.steering.CurrentVector(),
		TireStatus: tireDomain.BACKWARD,
	})
}

func (v *nexenTire) Stop() {
	v.TireStatus = tireDomain.STOP

	fmt.Println("정지")
}

func (v *nexenTire) Start() {
	v.TireStatus = tireDomain.START

	fmt.Println("부릉")
}
