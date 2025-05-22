package cmd

import (
	"context"
	"fmt"
	v1 "githubsearchservice/gen/github.com/akraturi/githubsearchservice/pkg/pb/v1"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Run the github search gRPC client",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("could not connect to server: %v", err)
		}
		defer func(conn *grpc.ClientConn) {
			err := conn.Close()
			if err != nil {

			}
		}(conn)

		client := v1.NewGithubSearchServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, err := client.Search(ctx, &v1.SearchRequest{Term: "kubernetes"})
		if err != nil {
			log.Fatalf("could not search term: %v", err)
		}
		fmt.Println("Response:", res.Results[0])
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
