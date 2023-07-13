package car

import (
	"designpattern_v1/v1/pattern/srp/domain"
	"designpattern_v1/v1/pattern/srp/maps"
	"fmt"
)

type volvo struct {
	maps *maps.Maps
}

func NewTire(maps *maps.Maps) domain.Tire {
	return &volvo{
		maps: maps,
	}
}

func (v *volvo) MoveForward() {
	coordinate := v.maps.Coordinate

	colLength := len(coordinate)

	if v.maps.CurrentX+1 < colLength {
		coordinate[v.maps.CurrentX][v.maps.CurrentY] = 0
		coordinate[v.maps.CurrentX+1][v.maps.CurrentY] = 1
		v.maps.CurrentX = v.maps.CurrentX + 1
	}

	if v.maps.CurrentX+1 >= colLength {
		fmt.Println("your car cannot move forward anymore!")
	}
}

func (v *volvo) MoveBackward() {
	coordinate := v.maps.Coordinate

	if v.maps.CurrentX-1 < 0 {
		fmt.Println("your car cannot move backward anymore!")
	}

	if v.maps.CurrentX-1 >= 0 {
		coordinate[v.maps.CurrentX][v.maps.CurrentY] = 0
		coordinate[v.maps.CurrentX-1][v.maps.CurrentY] = 1
		v.maps.CurrentX = v.maps.CurrentX - 1
	}

}

func (v *volvo) Stop() {
	fmt.Println("정지")
}

func (v *volvo) Start() {
	fmt.Println("부릉")
}
