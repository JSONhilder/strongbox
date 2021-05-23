package database

import (
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
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
	utils.PrintLogo()
	fmt.Println("Please enter your master password, strongbox does not store this.")
	utils.PrintWarn("It is up to you to remember this one, leave the rest to strongbox!")
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

func ExportDb(dst string) {
	var src string
	var newfile string
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	src = config.FilePath
	if runtime.GOOS == "windows" {
		newfile = dst + "\\strongbox"
	} else {
		newfile = dst + "/strongbox"
	}

	copyFile(src, newfile)

	utils.PrintSuccess("Strongbox successfully exported db file.")
}

func ImportDb(src string) {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	copyFile(src, config.FilePath)

	utils.PrintSuccess("Strongbox successfully imported db file.")
}

func copyFile(src, dst string) {
	from, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}
}
