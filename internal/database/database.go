package database

import (
	"encoding/gob"
	"log"
	"os"

	"github.com/JSONhilder/strongbox/internal/utils"
)

// Header is the main structure for the embedded "database"
type Header struct {
	Hk       string
	Sk       string
	Accounts []Account
}

// Account - Struct for users stored passwords
type Account struct {
	Name     string
	Username string
	Password string
	Url      string
}

var strongbox Header

// OpenDb -
func OpenDb(config *utils.Config) {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}
	// Checks if file exists first, if not create new one
	// @TODO: generate hash and salt from users master password
	if !fileExists(config.FilePath) {
		createStrongbox(config.FilePath)
	}

	f, err := os.Open(config.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	dec := gob.NewDecoder(f)

	// Send data to global strongbox
	if err := dec.Decode(&strongbox); err != nil {
		log.Fatal(err)
	}

}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func createStrongbox(filename string) {
	var accounts []Account
	header := Header{
		Hk:       "testHash",
		Sk:       "testSalt",
		Accounts: accounts,
	}

	f, err := os.Create(filename)
	if err != nil {
		return
	}
	defer f.Close()

	encoder := gob.NewEncoder(f)

	if err := encoder.Encode(header); err != nil {
		log.Println(err)
		return
	}
}

func writeData(updatedData Header) {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	f, err := os.Create(config.FilePath)
	if err != nil {
		return
	}
	defer f.Close()

	encoder := gob.NewEncoder(f)

	if err := encoder.Encode(updatedData); err != nil {
		log.Println(err)
		return
	}
}

// @TODO - func exportDb
