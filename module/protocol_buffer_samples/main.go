package main

import (
	"fmt"
	"protocol_buffer_samples/protos/model"
)

func main() {
	Read()
}

func Read() {
	protobuf := model.SearchRequest{
		Query:         "",
		PageNumber:    0,
		ResultPerPage: 0,
	}

	fmt.Println(protobuf)

}
