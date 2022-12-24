package casting

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TypeCasting(t *testing.T) {

	var badboys int = 1921

	var badboys2 float64 = float64(badboys)

	assert.Equal(t, 1921.0, badboys2)

	var badboys3 int64 = int64(badboys)

	assert.Equal(t, int64(1921), badboys3)

	var badboys4 uint = uint(badboys)

	assert.Equal(t, uint(1921), badboys4)

	var x int = 5
	var y int = 5
	var mul float32

	mul = float32(x) * float32(y)

	assert.Equal(t, float32(25), mul)

	// 아래의 코드는 golang은 강타입 시스템이기 때문에 에러가 발생함. 암묵적인 타입 캐스팅 불가
	// taking the required
	// data into variables
	// var x1 int = 19
	// var y1 float32 = 21
	// var mul1 float32

	//mul1 = x1 * y1

	// Displaying the result
	fmt.Printf("Multiplication = %f\n", mul)
}
