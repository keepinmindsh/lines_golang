package moving

import (
	"design_pattern/oop/app/maps"
	"design_pattern/oop/domain"
	movingDomain "design_pattern/oop/domain/moving"
)

type Moving struct {
	Backward             movingDomain.Moving
	Forward              movingDomain.Moving
	LeftForward          movingDomain.Moving
	RightForward         movingDomain.Moving
	LeftForwardDiagonal  movingDomain.Moving
	RightForwardDiagonal movingDomain.Moving
}

type xy struct {
	x int
	y int
}

var (
	UP                     = &xy{x: 0, y: 1}
	DOWN                   = &xy{x: 0, y: -1}
	LEFT_FORWARD           = &xy{x: -1, y: 0}
	RIGHT_FORWARD          = &xy{x: 1, y: 0}
	LEFT_FORWARD_DIAGONAL  = &xy{x: -1, y: 1}
	RIGHT_FORWARD_DIAGONAL = &xy{x: 1, y: 1}
)

func NewMoving(carName domain.CarName, maps *maps.Maps) *Moving {
	switch carName {
	case domain.TICO:
		return &Moving{
			Backward:             NewTicoMoveBackward(maps),
			Forward:              NewTicoMoveForward(maps),
			LeftForward:          NewTicoMoveLeftForward(maps),
			RightForward:         NewTicoMoveRightForward(maps),
			LeftForwardDiagonal:  NewTicoMoveLeftDiagonal(maps),
			RightForwardDiagonal: NewTicoMoveRightForward(maps),
		}
	default:
		return nil
	}
}
