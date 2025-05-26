package server

import (
	"fmt"
	v1 "githubsearchservice/gen/github.com/akraturi/githubsearchservice/pkg/pb/v1"
	"githubsearchservice/service/search"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	v1.GithubSearchServiceServer
	searcher search.Searcher
}

func New(searcher search.Searcher) Server {
	return Server{
		searcher: searcher,
	}
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	v1.RegisterGithubSearchServiceServer(server, s)

	log.Println("server listening on :50051")

	if err := server.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
