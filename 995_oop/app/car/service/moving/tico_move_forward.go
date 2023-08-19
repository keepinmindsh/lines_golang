package moving

import (
	"design_pattern/oop/app/maps"
	movingDomain "design_pattern/oop/domain/moving"
	"design_pattern/oop/internal/logger"
	"fmt"
)

type ticoMoveForward struct {
	maps *maps.Maps
}

func NewTicoMoveForward(maps *maps.Maps) movingDomain.Moving {
	return &ticoMoveForward{
		maps: maps,
	}
}

func (t *ticoMoveForward) Move() {
	t.maps.CurrentX = t.maps.CurrentX + UP.x
	t.maps.CurrentY = t.maps.CurrentY + UP.y

	logger.L.Info(fmt.Sprintf("Foward, Straight, Current Position - x : %v, y : %v \r\n", t.maps.CurrentX, t.maps.CurrentY))
}
