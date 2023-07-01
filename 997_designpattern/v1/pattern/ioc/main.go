package main

import (
	"designpattern_v1/v1/pattern/ioc/factory"
	"designpattern_v1/v1/pattern/ioc/make"
	"designpattern_v1/v1/pattern/ioc/product"
	"designpattern_v1/v1/pattern/ioc/receipt"
)

func main() {
	macComputer := receipt.MakeProductWithFactory[product.MacComputer](factory.ComputerFactory, make.Resource{
		Money:          50000,
		NumberOfPeople: 1,
		ProductCount:   50,
		Product:        make.MacBook,
	})

	macComputer.Start()

	macComputer.Compile()

	macComputer.Shutdown()
}
