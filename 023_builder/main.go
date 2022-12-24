package main

import (
	"023_builder/builder"
	"fmt"
)

func main() {
	assembly := builder.New()
	car := assembly.TopSpeed(1).Build()
	fmt.Println(car.Drive())
}
