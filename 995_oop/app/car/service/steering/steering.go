package steering

import (
	"design_pattern/oop/domain"
	"design_pattern/oop/internal/code"
	"fmt"
)

type steering struct {
	Vector domain.Vector
}

func (s *steering) RightDiagonal() {
	fmt.Println("우회전 설정")
	s.Vector = code.RightDiagonal
}

func (s *steering) LeftDiagonal() {
	fmt.Println("좌회전 설정")
	s.Vector = code.LeftDiagonal
}

func (s *steering) Strait() {
	fmt.Println("직진 설정")
	s.Vector = code.Straight
}

func (s *steering) CurrentVector() domain.Vector {
	return s.Vector
}

func NewSteering(vector domain.Vector) domain.Steering {
	return &steering{
		Vector: vector,
	}
}