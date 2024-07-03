package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pass",
	Short: "",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		log.Fatalf("Failed to execute command: %v", err.Error())
	}
}
