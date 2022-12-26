package sample

import (
	"bytes"
	"fmt"
)

func NestedLoop() {

	var b bytes.Buffer

	for i := 0; i < 100; i++ {
		for j := 0; j < 10; j++ {
			if i == j || (i%10) == j {
				b.WriteString("0-0 ")
			} else {
				b.WriteString(fmt.Sprintf(`%v - %v `, i, j))
			}
		}
		b.WriteString("\r\n")
	}

	fmt.Println(b.String())
}
