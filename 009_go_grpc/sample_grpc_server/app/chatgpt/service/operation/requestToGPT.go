package operation

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sample_grpc_server/domain"

	pb "github.com/keepinmindsh/go-lang-module/proto/gpt_sample"
	gogpt "github.com/sashabaranov/go-gpt3"
	"go.uber.org/zap"
)

type RequestGPT struct {
	pb.UnimplementedGPT3Server
	logger *zap.Logger
	client *gogpt.Client
}

func (r *RequestGPT) RequestToGPT(ctx context.Context, gptRequest domain.GPTRequest) *pb.GPT3Response {

	// GPT-3 API 호출을 위한 요청 생성
	stream, err := r.client.CreateCompletionStream(ctx, gogpt.CompletionRequest{
		Model:       gogpt.GPT3TextDavinci003,
		Prompt:      gptRequest.Prompt,
		MaxTokens:   700,
		Temperature: 0.5,
	})
	if err != nil {
		r.logger.Error("Error", zap.Error(err))
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
			fmt.Println("Code Sample - " + response.Choices[0].Text)
			contents += response.Choices[0].Text
		}
	}

	return &pb.GPT3Response{
		Text: contents,
	}
}

func NewRequestGPT(logger *zap.Logger, client *gogpt.Client) *RequestGPT {
	return &RequestGPT{
		logger: logger,
		client: client,
	}
}
