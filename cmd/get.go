package cmd

import (
	"github.com/JSONhilder/strongbox/internal/database"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get account by name from strongbox",
	Long:  `Get account with passed as argument name from strongbox`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		database.GetAccount(args[0])
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
