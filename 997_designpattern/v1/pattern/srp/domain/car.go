package domain

type (
	Tire interface {
		MoveForward()
		MoveBackward()
		Stop()
		Start()
	}
)
