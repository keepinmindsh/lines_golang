package tire

import (
	"design_pattern/oop/app/car/usecase/tire"
	"design_pattern/oop/domain"
)

type TireName string

const (
	KUMHO TireName = "Kumho"
	NEXEN TireName = "Nexen"
)

func NewTire(tireName TireName, steering domain.Steering, moving domain.Moving) domain.Tire {
	switch tireName {
	case KUMHO:
		return nil
	case NEXEN:
		return tire.NewNexenTire(steering, moving)
	}

	return nil
}
