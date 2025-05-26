package cmd

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"githubsearchservice/server"
	"githubsearchservice/service/search"
	"log"
	"os"
	"time"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the github search grpc server",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("error loading .env file")
		}

		githubApiClient := resty.New().
			SetBaseURL("https://api.github.com").
			SetTimeout(5*time.Second).
			SetRetryCount(3).
			AddRetryAfterErrorCondition().
			SetHeader("Content-Type", "application/json").
			SetHeader("Authorization",
				fmt.Sprintf("token %v", os.Getenv("GITHUB_TOKEN")))
		githubSearcher := search.NewGithubSearcher(githubApiClient)

		s := server.New(githubSearcher)
		if err := s.Run(); err != nil {
			return fmt.Errorf("failed to run server: %v", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
