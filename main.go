package main

import (
	"log"
	"net"

	"github.com/Alkestian/grpc-chat/grpc-chat/chat"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("error starting http server: %s", err)
	}
	
	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("error starting grpc server: %s", err)
	}
}