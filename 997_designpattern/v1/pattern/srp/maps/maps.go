package maps

type Maps struct {
	coordinate [][]int
}

func NewMaps(x, y int) Maps {
	var coordinate [][]int
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if i == 0 && j == 0 {
				coordinate[i][j] = 1
			}

			coordinate[i][j] = 0
		}
	}

	return Maps{coordinate: coordinate}
}
