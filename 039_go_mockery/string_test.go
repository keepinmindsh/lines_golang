package go_mockery

import (
	"github.com/stretchr/testify/assert"
	"go_mockery/mocks"
	"testing"
)

func TestString(t *testing.T) {
	stringer := mocks.Stringer{}
	stringer.EXPECT().String().Return("HighHopes")
	assert.Equal(t, "HighHopes", stringer.String())
}
