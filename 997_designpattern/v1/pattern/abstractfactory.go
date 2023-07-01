package pattern

import "fmt"

type Marine struct {
	health int
	shield int
}

func (marine Marine) Attack() {
	fmt.Println("Attack!")
}

func (marine *Marine) Damaged() {
	marine.shield -= 5

	if marine.shield < 0 {
		marine.health -= 10
	}

	fmt.Println("Damaged!")
}

type IUnitFactory interface {
	Attack()
	Damaged()
}

func GetMachineFactory(unitType string) (IUnitFactory, error) {
	switch unitType {
	case "Marine":
		return &Marine{health: 100, shield: 100}, nil
	}

	return nil, fmt.Errorf("wrong Unit Type")
}

func ExampleForAbstractFactory() {
	factory, err := GetMachineFactory("Machine")

	if err != nil {
		fmt.Errorf("error has been occured")
		return
	}

	factory.Damaged()
	factory.Attack()
}
