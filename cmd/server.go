package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"githubsearchservice/server"
	"githubsearchservice/service/search"
	"log"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the github search grpc server",
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("error loading .env file")
		}

		githubSearcher := search.NewGithubSearcher()

		s := server.New(githubSearcher)
		if err := s.Run(); err != nil {
			log.Fatal(fmt.Sprintf("failed to run server: %v", err))
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
