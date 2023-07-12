package main

import (
	"designpattern_v1/v1/pattern/srp/maps"
	"fmt"
)

func main() {
	newMaps := maps.NewMaps(5, 5)

	fmt.Println(newMaps)
}
