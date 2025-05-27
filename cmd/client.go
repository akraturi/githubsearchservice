package cmd

import (
	"context"
	"fmt"
	v1 "githubsearchservice/gen/github.com/akraturi/githubsearchservice/pkg/pb/v1"
	"log"
	"time"

	"google.golang.org/grpc/metadata"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	term           string //nolint:gochecknoglobals //spf13/cobra pattern
	user           string //nolint:gochecknoglobals //spf13/cobra pattern
	githubAPIToken string //nolint:gochecknoglobals //spf13/cobra pattern
)

//nolint:gochecknoglobals //spf13/cobra pattern
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Search for a term using github search api",
	RunE: func(_ *cobra.Command, _ []string) error {
		conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return fmt.Errorf("could not connect to server: %v", err)
		}
		defer func(conn *grpc.ClientConn) {
			err := conn.Close()
			if err != nil {
				log.Fatalf("failed to close client connection: %v", err)
			}
		}(conn)

		client := v1.NewGithubSearchServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		ctxWithMetaData := metadata.NewOutgoingContext(ctx, map[string][]string{
			"github_api_token": {githubAPIToken},
		})

		resp, err := client.Search(ctxWithMetaData, &v1.SearchRequest{Term: term, User: &user})
		if err != nil {
			return fmt.Errorf("error while doing grpc call: %v", err)
		}

		for _, result := range resp.GetResults() {
			log.Println(result)
		}

		return nil
	},
}

//nolint:gochecknoinits //spf13/cobra follows init pattern
func init() {
	clientCmd.Flags().StringVarP(&term, "term", "t", "", "term to search")
	clientCmd.Flags().StringVarP(&user, "user", "u", "", "user to search")
	clientCmd.Flags().StringVarP(&githubAPIToken, "github_api_token", "a", "",
		"github api token")
	rootCmd.AddCommand(clientCmd)
}
