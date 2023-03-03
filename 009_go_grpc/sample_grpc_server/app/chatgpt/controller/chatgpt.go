package controller

import (
	"context"
	"sample_grpc_server/app/chatgpt/service"
	"sample_grpc_server/domain"

	pb "github.com/keepinmindsh/go-lang-module/proto/gpt_sample"
	gogpt "github.com/sashabaranov/go-gpt3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type gptServer struct {
	pb.UnimplementedGPT3Server
	logger *zap.Logger
	client *gogpt.Client

	gtpService domain.GPTService
}

func (gpt *gptServer) GenerateText(ctx context.Context, gptRequest *pb.GPT3Request) (*pb.GPT3Response, error) {
	return gpt.gtpService.RequestToGPT(ctx, domain.GPTRequest{
		Model:       gptRequest.Model,
		Prompt:      gptRequest.Prompt,
		MaxTokens:   700,
		Temperature: 0.7,
	}), nil
}

func NewChatGPT(grpcServer *grpc.Server, logger *zap.Logger, client *gogpt.Client) {
	pb.RegisterGPT3Server(grpcServer, &gptServer{
		logger:     logger,
		client:     client,
		gtpService: service.NewGPTService(logger, client),
	})
}
