package tire

import (
	"design_pattern/oop/app/car/service/moving"
	tireDomain "design_pattern/oop/domain/tire"
	"fmt"
)

type nexenTire struct {
	moving     *moving.Moving
	TireStatus tireDomain.Status
}

func NewTicoTire(moving *moving.Moving) tireDomain.Tire {
	return &nexenTire{
		moving: moving,
	}
}

func (v *nexenTire) Forward() {
	v.TireStatus = tireDomain.FOWARD

	v.moving.Forward.Move()
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
