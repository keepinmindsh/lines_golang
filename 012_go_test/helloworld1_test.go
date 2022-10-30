package go_test

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go_test/model"
	"os"
	"time"
)

var _ = Describe("Good World/Bad World Test", func() {

	var dream *model.Dream

	BeforeEach(func() {
		fmt.Println("Run before starting each test module")

		dream = &model.Dream{
			Name:     "Good World",
			Deadline: time.Now(),
		}

		DeferCleanup(os.Setenv, "WEIGHT_UNITS", os.Getenv("WEIGHT_UNITS"))
	})

	AfterEach(func() {
		fmt.Println("Run before finished Test")
	})

	Describe("Dream is", func() {
		Context("making a good world", func() {
			It("Check dream's Name", func() {
				fmt.Printf("Dream is %s", dream.Name)
				Expect(dream.Name).To(Equal("Good World"))
			})
		})

		It("Check dream's Name", func() {
			fmt.Printf("Dream is %s", dream.Name)
			Expect(dream.Name).To(Equal("Bad World"))
		})
	})
})
