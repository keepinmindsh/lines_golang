package main

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/keepinmindsh/go-lang-module/proto/gpt_sample"
	gogpt "github.com/sashabaranov/go-gpt3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io"
)

// server is used to implement helloworld.GreeterServer.
type gpt_server struct {
	pb.UnimplementedGPT3Server
	logger *zap.Logger
}

func (gpt *gpt_server) GenerateText(ctx context.Context, gptRequest *pb.GPT3Request) (*pb.GPT3Response, error) {
	client := gogpt.NewClient("")

	// GPT-3 API 호출을 위한 요청 생성
	stream, err := client.CreateCompletionStream(ctx, gogpt.CompletionRequest{
		Model:       "text-davinci-002",
		Prompt:      "Hello, how are you today?",
		MaxTokens:   60,
		Temperature: 0.5,
	})
	if err != nil {
		gpt.logger.Error("Error", zap.Error(err))
	}

	defer stream.Close()

	var contents string

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("Stream finished")
			break
		}

		if err != nil {
			fmt.Printf("Stream error: %v\n", err)
		}

		if len(response.Choices) > 0 {
			contents += response.Choices[0].Text
		}
	}

	return &pb.GPT3Response{
		Text: contents,
	}, nil
}

func NewChatGPT(grpcServer *grpc.Server, logger *zap.Logger) {
	pb.RegisterGPT3Server(grpcServer, &gpt_server{logger: logger})
}
