package tire

import (
	"design_pattern/oop/app/car/service/moving"
	tierUcase "design_pattern/oop/app/car/usecase/tire"

	"design_pattern/oop/domain/tire"
)

type TireName string

const (
	KUMHO TireName = "Kumho"
	NEXEN TireName = "Nexen"
)

func NewTire(tireName TireName, moving *moving.Moving) tire.Tire {
	switch tireName {
	case KUMHO:
		return nil
	case NEXEN:
		return tierUcase.NewTicoTire(moving)
	}

	return nil
}
