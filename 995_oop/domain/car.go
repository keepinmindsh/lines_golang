package domain

type (
	Tire interface {
		MoveForward()
		MoveBackward()
		Stop()
		Start()
	}

	Steering interface {
		RightDiagonal()
		LeftDiagonal()
		Strait()
		CurrentVector() Vector
	}

	Car interface {
		Tire
		Steering
	}

	Moving interface {
		Move()
	}
)
