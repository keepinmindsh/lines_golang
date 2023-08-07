package moving

import (
	"design_pattern/oop/app/maps"
	"design_pattern/oop/domain"
)

type moving struct {
	maps maps.Maps
}

func NewMoving(carName domain.CarName, maps *maps.Maps) domain.Moving {
	switch carName {
	case domain.TICO:
		return NewTicoMoving(maps)
	default:
		return nil
	}
}
