package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"githubsearchservice/server"
	"log"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the github search gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		s := server.NewServer()
		if err := s.Run(); err != nil {
			log.Fatal(fmt.Sprintf("failed to run server: %v", err))
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
