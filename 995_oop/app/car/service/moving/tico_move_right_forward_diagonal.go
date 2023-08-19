package moving

import (
	"design_pattern/oop/app/maps"
	movingDomain "design_pattern/oop/domain/moving"
	"design_pattern/oop/internal/logger"
	"fmt"
)

type ticoMoveRightForwardDiagonal struct {
	maps *maps.Maps
}

func NewTicoMoveRightForwardDiagonal(maps *maps.Maps) movingDomain.Moving {
	return &ticoMoveRightForwardDiagonal{maps: maps}
}

func (t *ticoMoveRightForwardDiagonal) Move() {
	t.maps.CurrentX = t.maps.CurrentX + RIGHT_FORWARD_DIAGONAL.x
	t.maps.CurrentY = t.maps.CurrentY + RIGHT_FORWARD_DIAGONAL.y

	logger.L.Info(fmt.Sprintf("Foward, Straight, Current Position - x : %v, y : %v \r\n", t.maps.CurrentX, t.maps.CurrentY))
}
