package cmd

import (
	"fmt"
	"githubsearchservice/server"
	"githubsearchservice/service/search"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals //spf13/cobra pattern
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the github search grpc server",
	RunE: func(_ *cobra.Command, _ []string) error {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("error loading .env file")
		}

		githubAPIClient := resty.New().
			SetBaseURL("https://api.github.com").
			SetTimeout(5*time.Second).
			SetRetryCount(3).
			AddRetryAfterErrorCondition().
			SetHeader("Content-Type", "application/json")
		githubSearcher := search.NewGithubSearcher(githubAPIClient)

		s := server.New(githubSearcher)
		if err := s.Run(); err != nil {
			return fmt.Errorf("failed to run server: %v", err)
		}

		return nil
	},
}

//nolint:gochecknoinits //spf13/cobra follows init pattern
func init() {
	rootCmd.AddCommand(serverCmd)
}
