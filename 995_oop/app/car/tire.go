package car

import (
	"design_pattern/oop/app/car/component/tire"
	maps "design_pattern/oop/app/maps"
	"design_pattern/oop/domain"
)

type TireName string

const (
	KUMHO TireName = "Kumho"
	NEXEN TireName = "Nexen"
)

func NewTire(tireName TireName, maps *maps.Maps, steering domain.Steering) domain.Tire {
	switch tireName {
	case KUMHO:
		return nil
	case NEXEN:
		return tire.NewNexenTire(maps, steering)
	}

	return nil
}
