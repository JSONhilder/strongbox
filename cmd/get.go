package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var get = &cobra.Command{
	Use:   "get",
	Short: "Get account by name from stronbox",
	Long:  `Get account with passed as argument name from stronbox`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implented")
		if len(args) > 0 {
			fmt.Printf("First arg passed: %s", args[0])
			fmt.Println()
			fmt.Printf("All args: %v", args)
		}
	},
}

func init() {
	rootCmd.AddCommand(get)
}
