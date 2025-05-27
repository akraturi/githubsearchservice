package server

import (
	"fmt"
	v1 "githubsearchservice/gen/github.com/akraturi/githubsearchservice/pkg/pb/v1"
	"githubsearchservice/server/interceptors"
	"githubsearchservice/service/search"
	"log"
	"net"

	"google.golang.org/grpc"
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
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(interceptors.AuthInterceptor()),
	)
	v1.RegisterGithubSearchServiceServer(server, s)

	log.Println("server listening on localhost:50051")

	if err := server.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
