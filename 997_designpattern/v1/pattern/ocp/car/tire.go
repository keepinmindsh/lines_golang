package car

import (
	"designpattern_v1/v1/pattern/srp/domain"
	"designpattern_v1/v1/pattern/srp/maps"
)

type Tire string

const (
	KUMHO Tire = "Kumho"
	NEXEN Tire = "Nexen"
)

func NewTire(tire Tire, maps *maps.Maps) domain.Tire {
	switch tire {
	case KUMHO:
		return NewKumhoTire(maps)
	case NEXEN:
		return NewNexenTire(maps)
	}

	return nil
}
