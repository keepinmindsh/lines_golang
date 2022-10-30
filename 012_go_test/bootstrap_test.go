package go_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Test_BootStrap - 오직 자기 자신이 패키지에서만 현재는 동작하는 코드임.
func Test_BootStrap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Helloworld Suite")
}

// BeforeSuite - 모든 테스트 케이스 시작전에!
var _ = BeforeSuite(func() {
	fmt.Println("Run before start all test case")
})

// AfterSuite - 모든 테스트 케이스가 완료되었을 때!
var _ = AfterSuite(func() {
	fmt.Println("Run after start all test case")
})
