package product

type (
	Product interface {
		Apply(setting interface{})
		Get() interface{}
	}
	ProductType string
)

const (
	MacBook     ProductType = "MacBook"
	GalaxyBook  ProductType = "GalaxyBook"
	BMW         ProductType = "Bmw"
	AUDI        ProductType = "Audi"
	KIA         ProductType = "Kia"
	HYUNDAI     ProductType = "Hyundai"
	Iphone      ProductType = "IPhone"
	GalaxyPhone ProductType = "Galaxy"
)

func NewProduct(productType ProductType) Product {
	switch productType {
	case MacBook:
		return NewMacBookProduct()
	}
	return nil
}
