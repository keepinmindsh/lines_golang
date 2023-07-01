package pattern

import (
	"fmt"
	"testing"
)

func Test_CompositionSample(t *testing.T) {
	Process(&Object1{})
}

type Composition interface {
	Related()
}

func Process(composition Composition) {
	composition.Related()
}

type Object1 struct{}

func (o *Object1) Related() {
	fmt.Println("This is composition sample.")
}
