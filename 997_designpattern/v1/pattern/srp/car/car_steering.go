package car

import (
	"designpattern_v1/v1/pattern/srp/domain"
	"designpattern_v1/v1/pattern/srp/maps"
)

type Vector string

const (
	RightDiagonal Vector = "right_diagonal"
	LeftDiagonal  Vector = "left_diagonal"
	Straight      Vector = "straight"
)

type steering struct {
	Vector Vector
}

func NewSteering(maps *maps.Maps) domain.Steering {
	return &steering{}
}

func (s *steering) RightDiagonal() {
	s.Vector = RightDiagonal
}

func (s *steering) LeftDiagonal() {
	s.Vector = LeftDiagonal
}

func (s *steering) Straight() {
	s.Vector = Straight
}
