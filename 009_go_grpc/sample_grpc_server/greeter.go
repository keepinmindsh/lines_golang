package main

import (
	"context"
	pb "github.com/keepinmindsh/go-lang-module/proto/model"
	"google.golang.org/grpc"
	"log"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func NewGreeter(grpcServer *grpc.Server) {
	pb.RegisterGreeterServer(grpcServer, &server{})
}
