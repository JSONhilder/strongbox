package main

import (
	"os"

	"github.com/JSONhilder/strongbox/cmd"
	"github.com/JSONhilder/strongbox/internal/database"
	"github.com/JSONhilder/strongbox/internal/utils"
)

func init() {
	// Checks if db exists first, if not create new one
	if !database.FileExists(database.DatabaseDetails.Filepath) {
		database.CreateStrongbox(database.DatabaseDetails.Filepath)

		utils.PrintSuccess("Database has been created successfully.")
		os.Exit(0)
	}

	if os.Args[1] == "version" || os.Args[1] == "help" {
		utils.PrintLogo()
	}

	if os.Args[1] != "version" && os.Args[1] != "help" {
		database.OpenDb()
		if database.GainAccess() == true {
			return
		} else {
			os.Exit(0)
		}
	}
}

func main() {
	cmd.Execute()
}
