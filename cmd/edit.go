package cmd

import (
	"github.com/JSONhilder/strongbox/internal/database"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit specified account details from strongbox",
	Long:  `Edit specified account details from strongbox`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		database.EditAccount(args[0])
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
