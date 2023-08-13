package make

import (
	"designpattern_v1/v1/pattern/ioc/product"
	"fmt"
)

type PROCEDURE string

type MakeProcedure interface {
	GetDesign()
	MakeProduct()
	QA()
	ReturnFinalProduct() interface{}
}

const (
	COMPUTER PROCEDURE = "Computer"
	MOBILE   PROCEDURE = "Mobile"
	CAR      PROCEDURE = "Car"
)

type PRODUCT string

const (
	MacBook     PRODUCT = "MacBook"
	GalaxyBook  PRODUCT = "GalaxyBook"
	Iphone      PRODUCT = "IPhone"
	GalaxyPhone PRODUCT = "Galaxy"
	BMW         PRODUCT = "BMW"
	AUDI        PRODUCT = "Audi"
	HYUNDAI     PRODUCT = "Hyundai"
	KIA         PRODUCT = "KIA"
)

type Computer struct {
	resource Resource
	product  product.Product
}

type Resource struct {
	Money          int
	NumberOfPeople int
	ProductCount   int
	Product        PRODUCT
}

func (c Computer) GetDesign() {
	fmt.Println("Design")
}

func (c Computer) MakeProduct() {
	fmt.Println("Product")
}

func (c Computer) QA() {
	fmt.Println("QA")
}

func (c Computer) ReturnFinalProduct() interface{} {
	return c.product.Get()
}

func NewComputerProcedure(resource Resource) MakeProcedure {
	var realProduct product.Product

	switch resource.Product {
	case MacBook:
		realProduct = product.NewProduct(product.MacBook)
	case GalaxyBook:
		realProduct = product.NewProduct(product.GalaxyBook)
	default:
		realProduct = nil
	}

	return Computer{
		resource: resource,
		product:  realProduct,
	}
}
