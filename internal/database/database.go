package database

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/JSONhilder/strongbox/internal/crypt"
	"github.com/JSONhilder/strongbox/internal/utils"
	"golang.org/x/crypto/ssh/terminal"
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
var verification string

// OpenDb -
func OpenDb(config *utils.Config) {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
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

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateStrongbox(filename string) {
	header := buildHeader()

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

	header := Header{
		Hk:       string(hash),
		Sk:       salt,
		Accounts: accounts,
	}

	return header
}

func GainAccess() bool {
	fmt.Println("Please enter your master password:")
	password, _ := terminal.ReadPassword(int(syscall.Stdin))
	master := []byte(password)

	if crypt.VerifyHash(strongbox.Hk, master) == true {
		verification = string(password)
		return true
	}

	return false
}

func constructKey() string {
	pk := strongbox.Sk[(len(verification) + 1):]
	pk = (pk + "." + verification)
	return pk
}

// @TODO - func exportDb
