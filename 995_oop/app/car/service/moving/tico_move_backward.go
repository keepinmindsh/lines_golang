package moving

import (
	"design_pattern/oop/app/maps"
	movingDomain "design_pattern/oop/domain/moving"
	"design_pattern/oop/internal/logger"
	"fmt"
)

type ticoMoveBackward struct {
	maps *maps.Maps
}

func NewTicoMoveBackward(maps *maps.Maps) movingDomain.Moving {
	return &ticoMoveBackward{maps: maps}
}

func (t *ticoMoveBackward) Move() {
	t.maps.CurrentX = t.maps.CurrentX + DOWN.x
	t.maps.CurrentY = t.maps.CurrentY + DOWN.y

	logger.L.Info(fmt.Sprintf("Foward, Straight, Current Position - x : %v, y : %v \r\n", t.maps.CurrentX, t.maps.CurrentY))
}
