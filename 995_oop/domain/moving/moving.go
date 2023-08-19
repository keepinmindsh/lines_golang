package moving

import (
	steeringDomain "design_pattern/oop/domain/steering"
	tireDomain "design_pattern/oop/domain/tire"
)

type MoveOrder struct {
	Vector     steeringDomain.Vector
	TireStatus tireDomain.Status
}

type Moving interface {
	Move()
}
