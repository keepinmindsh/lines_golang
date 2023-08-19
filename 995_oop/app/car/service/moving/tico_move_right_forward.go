package moving

import (
	"design_pattern/oop/app/maps"
	movingDomain "design_pattern/oop/domain/moving"
	"design_pattern/oop/internal/logger"
	"fmt"
)

type ticoMoveRightForward struct {
	maps *maps.Maps
}

func NewTicoMoveRightForward(maps *maps.Maps) movingDomain.Moving {
	return &ticoMoveRightForward{maps: maps}
}

func (t *ticoMoveRightForward) Move() {
	t.maps.CurrentX = t.maps.CurrentX + RIGHT_FORWARD.x
	t.maps.CurrentY = t.maps.CurrentY + RIGHT_FORWARD.y

	logger.L.Info(fmt.Sprintf("Foward, Straight, Current Position - x : %v, y : %v \r\n", t.maps.CurrentX, t.maps.CurrentY))
}
