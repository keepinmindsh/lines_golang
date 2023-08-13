package moving

import (
	"design_pattern/oop/app/maps"
	"design_pattern/oop/domain"
	movingDomain "design_pattern/oop/domain/moving"
)

type moving struct {
	maps maps.Maps
}

func NewMoving(carName domain.CarName, maps *maps.Maps) movingDomain.Moving {
	switch carName {
	case domain.TICO:
		return NewTicoMoving(maps)
	default:
		return nil
	}
}
