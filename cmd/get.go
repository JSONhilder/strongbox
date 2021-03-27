package cmd

import (
	"fmt"

	"github.com/JSONhilder/strongbox/internal/database"
	"github.com/spf13/cobra"
)

var get = &cobra.Command{
	Use:   "get",
	Short: "Get account by name from stronbox",
	Long:  `Get account with passed as argument name from stronbox`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implented")
		database.GetAccount(args[0])
	},
}

func init() {
	rootCmd.AddCommand(get)
}
