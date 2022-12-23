package _21_floating_point_error

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
=== RUN   Test_Int

	floating_point_error_test.go:11:
	    	Error Trace:	/Users/lines/sources/02_linesgits/lines_golang/021_floating_point_error/floating_point_error_test.go:11
	    	Error:      	Not equal:
	    	            	expected: 1.2
	    	            	actual  : 1.2000000000000002
	    	Test:       	Test_Int

--- FAIL: Test_Int (0.00s)
*/
func Test_FloatingErrorTest(t *testing.T) {
	var l float64 = 0.2

	assert.Equal(t, 0.2*6, l*6)
}
