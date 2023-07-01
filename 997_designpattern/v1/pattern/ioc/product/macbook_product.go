package product

type MacBookProduct struct {
}

func (m MacBookProduct) Apply(setting interface{}) {
	//TODO implement me
	panic("implement me")
}

func (m MacBookProduct) Get() interface{} {
	return MacComputer{}
}

func NewMacBookProduct() Product {
	return MacBookProduct{}
}
