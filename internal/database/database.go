package database

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/JSONhilder/strongbox/internal/utils"
)

//go:embed strongbox.json
var f embed.FS

// Strongbox is the main structure for the embedded "database"
type Strongbox struct {
	Mhash    string    `json:"mhash"`
	Accounts []Account `json:"accounts"`
}

// Account - Struct for users stored passwords
type Account struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Url      string `json:"url"`
}

var database Strongbox

// OpenDb - Using go 1.16 open the embedded file, also setups db on first run
func OpenDb(config *utils.Config) {
	file, _ := f.ReadFile("strongbox.json")

	err := json.Unmarshal(file, &database)
	if err != nil {
		fmt.Println("error:", err)
	}

	// Make sure it has needed fields
	if database.Mhash == "" {
		fmt.Println("No database found, creating one now...")
		setupDb(database)
	}
}

func setupDb(database Strongbox) {
	database.Mhash = "Need to still create hash"
	updatedData, err := json.Marshal(database)
	if err != nil {
		log.Println("Failed to setup database: ", err.Error())
	}

	os.WriteFile("./internal/database/strongbox.json", updatedData, 644)
}

// @TODO - func exportDb
