package domain

type (
	Car interface {
		MoveForward()
		MoveBackward()
		Stop()
		StartDriving()
	}
)
