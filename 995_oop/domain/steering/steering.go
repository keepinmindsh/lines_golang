package steering

type Steering interface {
	RightDiagonal()
	LeftDiagonal()
	Strait()
	CurrentVector() Vector
}
