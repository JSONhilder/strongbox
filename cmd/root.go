package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "strongbox",
	Short: "Strongbox a simple password manager",
	Long:  `Strongbox a simple password manager, written in Go.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(0)
	}
}
