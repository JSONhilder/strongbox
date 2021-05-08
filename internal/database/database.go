package database

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"

	"github.com/JSONhilder/strongbox/internal/crypt"
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
	if !fileExists(config.FilePath) {
		// @TODO: generate hash and salt from users master password
		createStrongbox(config.FilePath)
	}

	f, err := os.Open(config.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	dec := gob.NewDecoder(f)
	// @TODO: before initialising global strongbox check the header has what it
	// needs

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
	header := buildHeader()

	// fmt.Println(header)

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

func buildHeader() Header {
	// @TODO: generate hash and salt from users master password
	var accounts []Account
	var masterPass string
	fmt.Println("Please enter your master password, strongbox does not keep this!")
	fmt.Println("It is up to you to remember this one...")
	fmt.Scan(&masterPass)

	hash, err := crypt.GenerateHash([]byte(masterPass))
	if err != nil {
		log.Fatal(err)
	}

	salt := crypt.GenerateKey(32)
	fmt.Println(len(salt))

	header := Header{
		Hk:       string(hash),
		Sk:       salt,
		Accounts: accounts,
	}

	return header
}

// @TODO - func exportDb
