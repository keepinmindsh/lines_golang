package tire

import (
	tierUcase "design_pattern/oop/app/car/usecase/tire"
	"design_pattern/oop/domain/moving"
	"design_pattern/oop/domain/steering"
	"design_pattern/oop/domain/tire"
)

type TireName string

const (
	KUMHO TireName = "Kumho"
	NEXEN TireName = "Nexen"
)

func NewTire(tireName TireName, steering steering.Steering, moving moving.Moving) tire.Tire {
	switch tireName {
	case KUMHO:
		return nil
	case NEXEN:
		return tierUcase.NewTicoTire(steering, moving)
	}

	return nil
}
