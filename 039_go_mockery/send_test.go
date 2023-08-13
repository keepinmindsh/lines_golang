package go_mockery

import (
	"github.com/stretchr/testify/assert"
	"go_mockery/mocks"
	"testing"
)

func TestName(t *testing.T) {
	expecter := mocks.SendFunc{}

	expecter.EXPECT().Execute("Data").Return(22, nil)

	execute, err := expecter.Execute("Data")
	assert.NoError(t, err, "Error has been occurred")

	assert.Equal(t, 22, execute)
}
