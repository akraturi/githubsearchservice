package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals //spf13/cobra pattern
var rootCmd = &cobra.Command{
	Use:   "githubsearchservice",
	Short: "github search grpc server and client",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("error while running command: %v", err)
	}
}
