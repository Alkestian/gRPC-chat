package chat

import "fmt"

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) JoinChat(joinReq *JoinChatRequest, joinServ ChatService_JoinChatServer) error {
	fmt.Printf("Received join chat request from %s\n", joinReq.ClientName)
	return nil
}