package factory

import (
	"designpattern_v1/v1/pattern/ioc/make"
	"designpattern_v1/v1/pattern/ioc/procedure"
)

type (
	Factory interface {
		GetProcedure(resource make.Resource) make.MakeProcedure
	}
)

func newComputerFactory() Factory {
	return &procedure.ComputerFactory{}
}

func newMobileFactory() Factory {
	return &procedure.MobileFactory{}
}

func newCarFactory() Factory {
	return &procedure.CarFactory{}
}

type OBJECT string

const (
	ComputerFactory OBJECT = "computerFactory"
	MobileFactory   OBJECT = "mobileFactory"
	CarFactory      OBJECT = "carFactory"
)

func NewMenuFactory(object OBJECT) Factory {
	switch object {
	case ComputerFactory:
		return newComputerFactory()
	case MobileFactory:
		return newMobileFactory()
	case CarFactory:
		return newCarFactory()
	}

	return nil
}
