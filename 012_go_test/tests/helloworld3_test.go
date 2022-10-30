package go_test

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Hello World Suite 1", func() {
	BeforeEach(func() {
		fmt.Println("Before Each")
	})

	Describe("Hello World Print", func() {
		Context("For golang, ", func() {
			It("It should be string template", func() {
				fmt.Println("Golang Print")
			})
		})

		Context("For Java", func() {
			It("It should be string primitive", func() {
				fmt.Println("Java Print")
			})
		})
	})
})
