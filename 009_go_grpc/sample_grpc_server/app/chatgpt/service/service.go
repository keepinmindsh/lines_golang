package service

import (
	"sample_grpc_server/app/chatgpt/service/operation"
	"sample_grpc_server/domain"

	gogpt "github.com/sashabaranov/go-gpt3"
	"go.uber.org/zap"
)

type gptService struct {
	*operation.RequestGPT
}

func NewGPTService(logger *zap.Logger, client *gogpt.Client) domain.GPTService {
	return &gptService{
		operation.NewRequestGPT(logger, client),
	}
}
