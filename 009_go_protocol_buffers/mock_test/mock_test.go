package mock_test

import (
	mock_model "009_go_protocol_buffers/proto/mock"
	"009_go_protocol_buffers/proto/model"
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCode(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)

	mockGreetingClient := mock_model.NewMockGreeterClient(ctrl)

	mockGreetingClient.EXPECT().SayHello(
		ctx,                // expect any value for first parameter
		gomock.Eq("Hello"), // expect any value for second parameter
	).Do(func(res model.HelloRequest) {
		fmt.Println(res.Name)
		assert.Equal(t, "Mocked RPC1", res.Name)
	}).Return(&model.HelloReply{Message: "Mocked RPC"}, nil).AnyTimes()
}
