package car

import (
	"designpattern_v1/v1/pattern/srp/domain"
	"designpattern_v1/v1/pattern/srp/maps"
)

type steering struct {
	maps *maps.Maps
}

func NewSteering(maps *maps.Maps) domain.Steering {
	return &steering{
		maps: maps,
	}
}

func (s *steering) RightDiagonal() {
	s.maps.IsRightDiagonal = true
	s.maps.IsRightDiagonal = false
	s.maps.IsStraight = false
}

func (s *steering) LeftDiagonal() {
	s.maps.IsRightDiagonal = false
	s.maps.IsRightDiagonal = true
	s.maps.IsStraight = false
}

func (s *steering) Straight() {
	s.maps.IsRightDiagonal = false
	s.maps.IsRightDiagonal = false
	s.maps.IsStraight = true
}
