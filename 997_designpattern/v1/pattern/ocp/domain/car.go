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
		Straight()
	}

	Car interface {
		Tire
		Steering
	}
)
