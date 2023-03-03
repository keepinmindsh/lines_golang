package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sample_grpc_server/app/chatgpt/controller"

	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	newServer := grpc.NewServer()

	// Set up Viper
	viper.SetConfigName("config")                                          // name of config file (without extension)
	viper.AddConfigPath("/Users/howard/sources/02_bong_git/lines_golang/") // search the current directory for the config file
	viper.SetConfigType("yaml")                                            // type of config file

	client := gogpt.NewClient(viper.GetString("gpt.client_key"))

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	NewGreeter(newServer)
	controller.NewChatGPT(newServer, logger, client)

	log.Printf("server listening at %v", lis.Addr())
	if err := newServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
