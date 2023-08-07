package factory

import (
	"design_pattern/oop/app/car/usecase"
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
		return usecase.NewNexenTire(maps, steering)
	}

	return nil
}
