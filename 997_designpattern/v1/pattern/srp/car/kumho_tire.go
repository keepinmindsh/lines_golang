package car

import (
	"designpattern_v1/v1/pattern/srp/domain"
	"designpattern_v1/v1/pattern/srp/maps"
	"fmt"
)

type kumhoTire struct {
	maps     *maps.Maps
	steering *steering
}

func NewKumhoTire(maps *maps.Maps) domain.Tire {
	return &kumhoTire{
		maps: maps,
	}
}

func (v *kumhoTire) MoveForward() {
	coordinate := v.maps.Coordinate

	colLength := len(coordinate)

	switch v.steering.Vector {
	case Straight:
		if v.maps.CurrentX+1 < colLength {
			coordinate[v.maps.CurrentX][v.maps.CurrentY] = 0
			coordinate[v.maps.CurrentX+1][v.maps.CurrentY] = 1
			v.maps.CurrentX = v.maps.CurrentX + 1
		}

		if v.maps.CurrentX+1 >= colLength {
			fmt.Println("your car cannot move forward anymore!")
		}

		break
	case RightDiagonal:
		break
	case LeftDiagonal:
		break
	}
}

func (v *kumhoTire) MoveBackward() {
	coordinate := v.maps.Coordinate

	switch v.steering.Vector {
	case Straight:
		if v.maps.CurrentX-1 < 0 {
			fmt.Println("your car cannot move backward anymore!")
		}

		if v.maps.CurrentX-1 >= 0 {
			coordinate[v.maps.CurrentX][v.maps.CurrentY] = 0
			coordinate[v.maps.CurrentX-1][v.maps.CurrentY] = 1
			v.maps.CurrentX = v.maps.CurrentX - 1
		}
		break
	case RightDiagonal:
		break
	case LeftDiagonal:
		break
	}

}

func (v *kumhoTire) Stop() {
	fmt.Println("정지")
}

func (v *kumhoTire) Start() {
	fmt.Println("부릉")
}
