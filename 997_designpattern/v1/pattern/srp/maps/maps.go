package maps

type Maps struct {
	Coordinate      [][]int
	CurrentX        int
	CurrentY        int
	IsRightDiagonal bool
	IsLeftDiagonal  bool
	IsStraight      bool
}

func NewMaps(x, y int) Maps {
	coordinate := make([][]int, 5)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			coordinate[i] = make([]int, 5)
			if i == 0 && j == 0 {
				coordinate[i][j] = 1
			}

			coordinate[i][j] = 0
		}
	}

	return Maps{
		Coordinate:      coordinate,
		CurrentX:        0,
		CurrentY:        0,
		IsRightDiagonal: false,
		IsLeftDiagonal:  false,
		IsStraight:      false,
	}
}
