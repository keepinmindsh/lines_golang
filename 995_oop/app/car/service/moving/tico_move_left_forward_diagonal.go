package moving

import (
	"design_pattern/oop/app/maps"
	movingDomain "design_pattern/oop/domain/moving"
	"design_pattern/oop/internal/logger"
	"fmt"
)

type ticoMoveLeftForwardDiagonal struct {
	maps *maps.Maps
}

func NewTicoMoveLeftDiagonal(maps *maps.Maps) movingDomain.Moving {
	return &ticoMoveLeftForwardDiagonal{maps: maps}
}

func (t *ticoMoveLeftForwardDiagonal) Move() {
	t.maps.CurrentX = t.maps.CurrentX + LEFT_FORWARD_DIAGONAL.x
	t.maps.CurrentY = t.maps.CurrentY + LEFT_FORWARD_DIAGONAL.y

	logger.L.Info(fmt.Sprintf("Foward, Straight, Current Position - x : %v, y : %v \r\n", t.maps.CurrentX, t.maps.CurrentY))
}
