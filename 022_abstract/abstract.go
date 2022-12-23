package main

import "fmt"

type AbstractModule interface {
	GetAbstract(string) string
}

type abstract struct {
	value string
}

func NewAbstract() AbstractModule {
	return &abstract{
		value: "Hello World",
	}
}

func (a *abstract) GetAbstract(value string) string {
	return a.value
}

func main() {

	newAbstract := NewAbstract()

	fmt.Println(fmt.Sprintf("%s", newAbstract.GetAbstract("Hello")))
}
