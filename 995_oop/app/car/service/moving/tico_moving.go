package moving

import (
	"design_pattern/oop/app/maps"
	movingDomain "design_pattern/oop/domain/moving"
	"fmt"
)

type ticoMoving struct {
	maps *maps.Maps
}

func NewTicoMoving(maps *maps.Maps) movingDomain.Moving {
	return &ticoMoving{
		maps: maps,
	}
}

func (t *ticoMoving) Move(moveOrder movingDomain.MoveOrder) {
	colLength := len(t.maps.Coordinate)

	if t.maps.CurrentX+1 < colLength {
		t.maps.Coordinate[t.maps.CurrentX][t.maps.CurrentY] = 0
		t.maps.Coordinate[t.maps.CurrentX+1][t.maps.CurrentY] = 1
		t.maps.CurrentX = t.maps.CurrentX + 1
	}

	if t.maps.CurrentX+1 >= colLength {
		fmt.Println("your car cannot move forward anymore!")
	}

	fmt.Printf("Foward, Straight, Current Position - x : %v, y : %v \r\n", t.maps.CurrentX, t.maps.CurrentY)
}
