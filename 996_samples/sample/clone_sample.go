package sample

import (
	"fmt"
)

func CloneTest() {

	cloneObj := NewClone()

	clone := cloneObj.Clone()

	fmt.Printf("%v \r\n", &cloneObj)
	fmt.Printf("%v \r\n", &clone)

	ints, ok := clone.([]int)

	if ok {
		for _, value := range ints {
			fmt.Println(value)
		}
	} else {
		fmt.Println(fmt.Errorf("type Error : %s", "type"))
	}

}

type Cloneable interface {
	Clone() interface{}
}

type CloneObject struct {
	cloneObject []int
}

func (c *CloneObject) Clone() interface{} {
	dest := make([]int, len(c.cloneObject))

	if n := copy(dest, c.cloneObject); n != len(c.cloneObject) {
		fmt.Println("복사가 덜 됐습니다.")
	}

	return dest
}

func NewClone() Cloneable {
	return &CloneObject{
		cloneObject: []int{1, 2, 3, 4, 5, 6},
	}
}
