package builder

import "strconv"

type Color string
type Make string
type Model string

const (
	BLUE Color = "blue"
	RED        = "red"
)

type car struct {
	topSpeed int
	color    Color
}
type Car interface {
	Drive() string
	Stop() string
}

type CarBuilder interface {
	TopSpeed(int) CarBuilder
	Paint(Color) CarBuilder
	Build() Car
}

type carBuilder struct {
	speedOption int
	color       Color
}

func (cb *carBuilder) TopSpeed(speed int) CarBuilder {
	cb.speedOption = speed
	return cb
}

func (cb *carBuilder) Paint(color Color) CarBuilder {
	cb.color = color
	return cb
}

func (cb *carBuilder) Build() Car {
	return &car{
		topSpeed: cb.speedOption,
		color:    cb.color,
	}
}

func New() CarBuilder {
	return &carBuilder{}
}

func (c *car) Drive() string {
	return "Driving at speed: " + strconv.Itoa(c.topSpeed)
}

func (c *car) Stop() string {
	return "Stopping a " + string(c.color) + " car"
}
