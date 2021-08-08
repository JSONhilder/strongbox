package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/JSONhilder/strongbox/cmd"
	"github.com/JSONhilder/strongbox/internal/database"
	"github.com/JSONhilder/strongbox/internal/utils"
)

func init() {
	// Set database path
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	database.DatabaseDetails.Filepath = strings.Replace(ex, "strongbox.exe", "strongbox_db", 1)

	// Checks if db exists first, if not create new one
	if !database.FileExists(database.DatabaseDetails.Filepath) {
		database.CreateStrongbox(database.DatabaseDetails.Filepath)
		utils.PrintSuccess("Database has been created successfully.")
		os.Exit(0)
	}

	if len(os.Args) == 1 {
		utils.PrintLogo()
		fmt.Println("No command found, for help type strongbox help")
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
