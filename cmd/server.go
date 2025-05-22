package cmd

import (
	"github.com/spf13/cobra"
	v1 "githubsearchservice/gen/github.com/akraturi/githubsearchservice/pkg/pb/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	v1.UnimplementedGithubSearchServiceServer
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the github search gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		v1.RegisterGithubSearchServiceServer(s, Server{})

		log.Println("server listening on :50051")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
