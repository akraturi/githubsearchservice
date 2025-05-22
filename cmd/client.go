package cmd

import (
	"context"
	v1 "githubsearchservice/gen/github.com/akraturi/githubsearchservice/pkg/pb/v1"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var term string
var user string

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
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		resp, err := client.Search(ctx, &v1.SearchRequest{Term: term, User: &user})
		if err != nil {
			log.Fatalf("could not search term: %v", err)
		}

		for _, result := range resp.GetResults() {
			log.Println(result)
		}
	},
}

func init() {
	clientCmd.Flags().StringVarP(&term, "term", "t", "", "term to search")
	clientCmd.Flags().StringVarP(&user, "user", "u", "", "user to search")
	rootCmd.AddCommand(clientCmd)
}
