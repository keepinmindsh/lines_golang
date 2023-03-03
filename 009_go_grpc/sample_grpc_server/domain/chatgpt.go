package domain

import (
	"context"

	pb "github.com/keepinmindsh/go-lang-module/proto/gpt_sample"
)

type GPTService interface {
	RequestToGPT(ctx context.Context, gptRequest GPTRequest) *pb.GPT3Response
}

type GPTRequest struct {
	Model       string
	Prompt      string
	MaxTokens   int64
	Temperature float64
}
