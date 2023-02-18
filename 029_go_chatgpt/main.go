package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "github.com/keepinmindsh/go-lang-module/proto/gpt_sample"
	gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {
	// OpenAI API Key
	apiKey := "YOUR_API_KEY"

	// gRPC 서버에 연결
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// gRPC 클라이언트 생성
	client := pb.NewGPT3Client(conn)

	// OpenAI 클라이언트 생성
	openaiClient := gogpt.NewClient(apiKey)

	// gRPC 요청 생성
	gpt3Req := &pb.GPT3Request{
		Model:  req.Model,
		Prompt: req.Prompt,
	}

	// gRPC 서버로 요청 전송
	resp, err := client.GenerateText(context.Background(), gpt3Req)
	if err != nil {
		log.Fatalf("Failed to call GenerateText RPC: %v", err)
	}

	// OpenAI API 응답 출력
	log.Printf("Generated text: %s", resp.Text)

}
