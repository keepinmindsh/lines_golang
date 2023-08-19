package tire

import (
	"design_pattern/oop/app/car/service/moving"
	"design_pattern/oop/domain/steering"
	tireDomain "design_pattern/oop/domain/tire"
	"fmt"
)

type ticoTire struct {
	steering   steering.Steering
	moving     moving.Moving
	TireStatus tireDomain.Status
}

func NewTicoTire(steering steering.Steering, moving moving.Moving) tireDomain.Tire {
	return &ticoTire{
		steering: steering,
		moving:   moving,
	}
}

func (v *ticoTire) Forward() {
	v.TireStatus = tireDomain.FOWARD

	v.moving.Forward.Move()
}

func (v *ticoTire) Backward() {
	v.TireStatus = tireDomain.BACKWARD

	v.moving.Backward.Move()
}

func (v *ticoTire) Stop() {
	v.TireStatus = tireDomain.STOP

	fmt.Println("정지")
}

func (v *ticoTire) Start() {
	v.TireStatus = tireDomain.START

	fmt.Println("부릉")
}
