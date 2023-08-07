package tire

import (
	"design_pattern/oop/domain"
	"design_pattern/oop/internal/code"
	"fmt"
)

type nexenTire struct {
	steering domain.Steering
	moving   domain.Moving
}

func NewNexenTire(steering domain.Steering, moving domain.Moving) domain.Tire {
	return &nexenTire{
		steering: steering,
		moving:   moving,
	}
}

func (v *nexenTire) MoveForward() {
	switch v.steering.CurrentVector() {
	case code.Straight:
		v.moving.Move()
	case code.RightDiagonal:
		v.moving.Move()
	case code.LeftDiagonal:
		v.moving.Move()
	}
}

func (v *nexenTire) MoveBackward() {
	switch v.steering.CurrentVector() {
	case code.Straight:
	case code.RightDiagonal:
	case code.LeftDiagonal:
	}

}

func (v *nexenTire) Stop() {
	fmt.Println("정지")
}

func (v *nexenTire) Start() {
	fmt.Println("부릉")
}
