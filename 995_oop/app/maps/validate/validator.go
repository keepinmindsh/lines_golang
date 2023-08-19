package validate

import (
	"design_pattern/oop/app/maps"
	domainMaps "design_pattern/oop/domain/maps"
)

type validate struct {
	maps *maps.Maps
}

func NewValidater(maps *maps.Maps) domainMaps.MapValidate {
	return &validate{maps: maps}
}

func (v validate) Valid() bool {
	//TODO implement me
	panic("implement me")
}
