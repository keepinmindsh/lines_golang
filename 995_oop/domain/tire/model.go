package tire

type Status string

const (
	FOWARD   Status = "+"
	BACKWARD Status = "-"
	STOP     Status = "Stop"
	START    Status = "Start"
)
