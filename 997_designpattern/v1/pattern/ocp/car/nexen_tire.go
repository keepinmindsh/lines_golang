package car

import (
	"designpattern_v1/v1/pattern/srp/domain"
	"designpattern_v1/v1/pattern/srp/maps"
	"fmt"
)

type nexenTire struct {
	maps     *maps.Maps
	steering *steering
}

func NewNexenTire(maps *maps.Maps) domain.Tire {
	return &nexenTire{
		maps: maps,
	}
}

func (v *nexenTire) MoveForward() {
	coordinate := v.maps.Coordinate

	colLength := len(coordinate)
	rowLength := len(coordinate[0])

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

		if v.maps.CurrentX+1 < colLength && v.maps.CurrentY+1 < rowLength {
			coordinate[v.maps.CurrentX][v.maps.CurrentY] = 0
			coordinate[v.maps.CurrentX+1][v.maps.CurrentY+1] = 1
			v.maps.CurrentX = v.maps.CurrentX + 1
			v.maps.CurrentY = v.maps.CurrentY + 1
		}

		if v.maps.CurrentX >= colLength {
			fmt.Println("your car cannot move right diagonal forward anymore!")
		}

		if v.maps.CurrentY >= rowLength {
			fmt.Println("your car cannot move right diagonal forward anymore!")
		}

		break
	case LeftDiagonal:

		if v.maps.CurrentX+1 < colLength && v.maps.CurrentY+1 < rowLength {
			coordinate[v.maps.CurrentX][v.maps.CurrentY] = 0
			coordinate[v.maps.CurrentX+1][v.maps.CurrentY-1] = 1
			v.maps.CurrentX = v.maps.CurrentX + 1
			v.maps.CurrentY = v.maps.CurrentY - 1
		}

		if v.maps.CurrentX >= colLength {
			fmt.Println("your car cannot move left diagonal forward anymore!")
		}

		if v.maps.CurrentY <= 0 {
			fmt.Println("your car cannot move left diagonal forward anymore!")
		}

		break
	}
}

func (v *nexenTire) MoveBackward() {
	coordinate := v.maps.Coordinate

	rowLength := len(coordinate[0])

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

		if v.maps.CurrentX-1 < 0 {
			fmt.Println("your car cannot move right backward anymore!")
		}

		if v.maps.CurrentX+1 >= rowLength {
			fmt.Println("your car cannot move right backward anymore!")
		}

		if v.maps.CurrentX-1 >= 0 {
			coordinate[v.maps.CurrentX][v.maps.CurrentY] = 0
			coordinate[v.maps.CurrentX-1][v.maps.CurrentY+1] = 1
			v.maps.CurrentX = v.maps.CurrentX - 1
			v.maps.CurrentY = v.maps.CurrentY + 1
		}

		break
	case LeftDiagonal:

		if v.maps.CurrentX-1 < 0 {
			fmt.Println("your car cannot move right backward anymore!")
		}

		if v.maps.CurrentX-1 < 0 {
			fmt.Println("your car cannot move right backward anymore!")
		}

		if v.maps.CurrentX-1 >= 0 {
			coordinate[v.maps.CurrentX][v.maps.CurrentY] = 0
			coordinate[v.maps.CurrentX-1][v.maps.CurrentY+1] = 1
			v.maps.CurrentX = v.maps.CurrentX - 1
			v.maps.CurrentY = v.maps.CurrentY - 1
		}

		break
	}

}

func (v *nexenTire) Stop() {
	fmt.Println("정지")
}

func (v *nexenTire) Start() {
	fmt.Println("부릉")
}
