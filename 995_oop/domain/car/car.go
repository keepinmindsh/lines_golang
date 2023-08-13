package car

import (
	"design_pattern/oop/domain/steering"
	"design_pattern/oop/domain/tire"
)

type Car interface {
	tire.Tire
	steering.Steering
}
