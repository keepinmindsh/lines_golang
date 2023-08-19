package moving

import (
	"design_pattern/oop/app/maps"
	movingDomain "design_pattern/oop/domain/moving"
	"design_pattern/oop/internal/logger"
	"fmt"
)

type ticoMoveLeftForward struct {
	maps *maps.Maps
}

func NewTicoMoveLeftForward(maps *maps.Maps) movingDomain.Moving {
	return &ticoMoveLeftForward{maps: maps}
}

func (t *ticoMoveLeftForward) Move() {
	t.maps.CurrentX = t.maps.CurrentX + LEFT_FORWARD.x
	t.maps.CurrentY = t.maps.CurrentY + LEFT_FORWARD.y

	logger.L.Info(fmt.Sprintf("Foward, Straight, Current Position - x : %v, y : %v \r\n", t.maps.CurrentX, t.maps.CurrentY))
}
