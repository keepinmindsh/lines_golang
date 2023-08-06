package tire

import (
	maps "design_pattern/oop/app/maps"
	"design_pattern/oop/domain"
	"design_pattern/oop/internal/code"
	"fmt"
)

type nexenTire struct {
	maps     *maps.Maps
	steering domain.Steering
}

func NewNexenTire(maps *maps.Maps, steering domain.Steering) domain.Tire {
	return &nexenTire{
		maps:     maps,
		steering: steering,
	}
}

func (v *nexenTire) MoveForward() {
	coordinate := v.maps.Coordinate

	colLength := len(coordinate)

	switch v.steering.CurrentVector() {
	case code.Straight:
		if v.maps.CurrentX+1 < colLength {
			coordinate[v.maps.CurrentX][v.maps.CurrentY] = 0
			coordinate[v.maps.CurrentX+1][v.maps.CurrentY] = 1
			v.maps.CurrentX = v.maps.CurrentX + 1
		}

		if v.maps.CurrentX+1 >= colLength {
			fmt.Println("your car cannot move forward anymore!")
		}

		fmt.Printf("Straight, Current Position - x : %v, y : %v \r\n", v.maps.CurrentX, v.maps.CurrentY)
	case code.RightDiagonal:
		fmt.Printf("After RightDiagonal, Current Position - x : %v, y : %v \r\n", v.maps.CurrentX, v.maps.CurrentY)
	case code.LeftDiagonal:
		fmt.Printf("After LeftDiagonal, Current Position - x : %v, y : %v \r\n", v.maps.CurrentX, v.maps.CurrentY)
	}
}

func (v *nexenTire) MoveBackward() {
	coordinate := v.maps.Coordinate

	switch v.steering.CurrentVector() {
	case code.Straight:
		if v.maps.CurrentX-1 < 0 {
			fmt.Println("your car cannot move backward anymore!")
		}

		if v.maps.CurrentX-1 >= 0 {
			coordinate[v.maps.CurrentX][v.maps.CurrentY] = 0
			coordinate[v.maps.CurrentX-1][v.maps.CurrentY] = 1
			v.maps.CurrentX = v.maps.CurrentX - 1
		}

		fmt.Printf("Straight, Current Position - x : %v, y : %v \r\n", v.maps.CurrentX, v.maps.CurrentY)
	case code.RightDiagonal:
		fmt.Printf("After RightDiagonal, Current Position - x : %v, y : %v \r\n", v.maps.CurrentX, v.maps.CurrentY)
	case code.LeftDiagonal:
		fmt.Printf("After LeftDiagonal, Current Position - x : %v, y : %v \r\n", v.maps.CurrentX, v.maps.CurrentY)
	}

}

func (v *nexenTire) Stop() {
	fmt.Println("정지")
}

func (v *nexenTire) Start() {
	fmt.Println("부릉")
}
