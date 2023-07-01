package receipt

import (
	"designpattern_v1/v1/pattern/ioc/factory"
	"designpattern_v1/v1/pattern/ioc/make"
)

func MakeProductWithFactory[T any](object factory.OBJECT, resource make.Resource) T {
	// Interface 를 반환받는다.
	factory := factory.NewMenuFactory(object)

	makeProduct := factory.GetProcedure(resource)

	makeProduct.GetDesign()

	makeProduct.MakeProduct()

	makeProduct.GetDesign()

	return makeProduct.ReturnFinalProduct().(T)
}
