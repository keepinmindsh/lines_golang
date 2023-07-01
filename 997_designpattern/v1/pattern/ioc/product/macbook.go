package product

type MacComputer struct {
}

type MacBookProduct struct {
}

func (m MacBookProduct) Apply(setting interface{}) {
	//TODO implement me
	panic("implement me")
}

func (m MacBookProduct) Get() interface{} {
	return MacComputer{}
}

func (mac MacComputer) Start() {

}

func (mac MacComputer) Compile() {

}

func (mac MacComputer) Shutdown() {

}

func NewMacBookProduct() Product {
	return MacBookProduct{}
}
