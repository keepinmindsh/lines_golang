package maps

type Maps struct {
	Coordinate [][]int
	CurrentX   int
	CurrentY   int
}

const (
	startPoint     int = 1
	startPositionX int = 0
	startPositionY int = 0
	positionInit   int = 0
)

func NewMaps(x, y int) Maps {
	coordinate := make([][]int, x)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			coordinate[i] = make([]int, y)
			if isStartPoint(i, j) {
				coordinate[i][j] = startPoint
			}

			coordinate[i][j] = positionInit
		}
	}

	return Maps{
		Coordinate: coordinate,
		CurrentX:   startPositionX,
		CurrentY:   startPositionY,
	}
}

func isStartPoint(i int, j int) bool {
	return i == startPositionX && j == startPositionY
}
