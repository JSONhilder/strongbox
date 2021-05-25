package main

import (
	"log"
	"os"

	"github.com/JSONhilder/strongbox/cmd"
	"github.com/JSONhilder/strongbox/internal/database"
	"github.com/JSONhilder/strongbox/internal/utils"
)

func init() {
	// Setup/Run anything needed before cli executes
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Println("No config file found.")
		os.Exit(-1)
	}

	// Checks if db exists first, if not create new one
	if !database.FileExists(config.FilePath) {
		database.CreateStrongbox(config.FilePath)

		utils.PrintSuccess("Database has been created successfully.")
		os.Exit(0)
	}

	if os.Args[1] == "version" || os.Args[1] == "help" {
		utils.PrintLogo()
	}

	if os.Args[1] != "version" && os.Args[1] != "help" {
		database.OpenDb(config)
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
