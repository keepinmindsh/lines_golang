package tire

import (
	"design_pattern/oop/app/car/service/moving"
	tierUcase "design_pattern/oop/app/car/usecase/tire"
	steeringDomain "design_pattern/oop/domain/steering"

	"design_pattern/oop/domain/tire"
)

type TireName string

const (
	KUMHO TireName = "Kumho"
	NEXEN TireName = "Nexen"
)

func NewTire(tireName TireName, moving *moving.Moving, steering *steeringDomain.Steering) tire.Tire {
	switch tireName {
	case KUMHO:
		return nil
	case NEXEN:
		return tierUcase.NewNexenTire(moving, steering)
	}

	return nil
}
