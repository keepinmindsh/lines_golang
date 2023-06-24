package designpattern

import "fmt"

func PatternAddicted() {
	hello := NewFactoryHello(PRINT)

	err := hello.Do("hello")
	if err != nil {
		panic(err)
	}
}

type (
	action interface {
		Do(value any) error
	}
)
type Print struct{}
type ACTION string

const (
	PRINT ACTION = "Print"
)

func NewFactoryHello(action ACTION) action {
	switch action {
	case PRINT:
		return NewPrintAction()
	default:
		return nil
	}
}

func NewPrintAction() action {
	return Print{}
}

func (p Print) Do(value any) error {
	s, ok := value.(string)
	if ok {
		fmt.Println(s)
		return nil
	} else {
		err := fmt.Errorf("error type is wrong")
		return err
	}
}
