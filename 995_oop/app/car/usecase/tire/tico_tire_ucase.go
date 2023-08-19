package tire

import (
	"design_pattern/oop/app/car/service/moving"
	tireDomain "design_pattern/oop/domain/tire"
	"fmt"
)

type ticoTire struct {
	moving     *moving.Moving
	TireStatus tireDomain.Status
}

func NewTicoTire(moving *moving.Moving) tireDomain.Tire {
	return &ticoTire{
		moving: moving,
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
