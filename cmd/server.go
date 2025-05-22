package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"githubsearchservice/server"
	"githubsearchservice/service/githubservice"
	"log"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the github search grpc server",
	Run: func(cmd *cobra.Command, args []string) {
		githubService := githubservice.New()

		s := server.New(githubService)
		if err := s.Run(); err != nil {
			log.Fatal(fmt.Sprintf("failed to run server: %v", err))
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
