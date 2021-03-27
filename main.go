package main

import (
	"log"

	"github.com/JSONhilder/strongbox/cmd"
	"github.com/JSONhilder/strongbox/internal/database"
	"github.com/JSONhilder/strongbox/internal/utils"
)

func init() {
	// Setup/Run anything needed before cli executes
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Println("No config file found.")
	}

	database.OpenDb(config)
}

func main() {
	cmd.Execute()
}
