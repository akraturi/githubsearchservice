package server

import (
	"errors"
	"fmt"
	v1 "githubsearchservice/gen/github.com/akraturi/githubsearchservice/pkg/pb/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	v1.GithubSearchServiceServer
}

func NewServer() Server {
	return Server{}
}

func (server *Server) Run() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return errors.New(fmt.Sprintf("failed to listen: %v", err))
	}

	s := grpc.NewServer()
	v1.RegisterGithubSearchServiceServer(s, server)

	log.Println("server listening on :50051")

	if err := s.Serve(lis); err != nil {
		return errors.New(fmt.Sprintf("failed to serve: %v", err))
	}

	return nil
}
