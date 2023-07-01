package procedure

import (
	"designpattern_v1/v1/pattern/ioc/make"
)

type (
	ComputerFactory struct{}

	MobileFactory struct{}

	CarFactory struct{}
)

func (f CarFactory) GetProcedure(resource make.Resource) make.MakeProcedure {
	switch resource.Product {
	case make.BMW:
		return nil
	case make.AUDI:
		return nil
	case make.KIA:
		return nil
	case make.HYUNDAI:
		return nil
	}

	return nil
}

func (f MobileFactory) GetProcedure(resource make.Resource) make.MakeProcedure {
	switch resource.Product {
	case make.Iphone:
		return nil
	case make.GalaxyPhone:
		return nil
	}

	return nil
}

func (f ComputerFactory) GetProcedure(resource make.Resource) make.MakeProcedure {
	switch resource.Product {
	case make.MacBook:
		return make.NewComputerProcedure(resource)
	case make.GalaxyBook:
		return make.NewComputerProcedure(resource)
	}

	return nil
}
