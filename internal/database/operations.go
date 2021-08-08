package database

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/JSONhilder/strongbox/internal/crypt"
	"github.com/JSONhilder/strongbox/internal/utils"
)

func ListAccounts() {
	key := constructKey()
	if strongbox.Accounts == nil {
		utils.PrintError("No accounts in database")
		fmt.Println("Create one with the save command")
		return
	}

	fmt.Println("Stronbox Accounts")
	fmt.Println("-----------------")
	for index, acc := range strongbox.Accounts {
		fmt.Printf("%d: %s\n", index, crypt.DecryptKey(acc.Name, key))
	}
}

func GetAccount(name string) {
	// text to clipboarb win10 echo mytext | clip
	// text to clipboarb macOS echo mytext | pbcopy
	// Need to figure out a work around for linux
	index, found, key := doesAccountExist(name)

	if found == true {
		var arg string
		account := strongbox.Accounts[index]
		fmt.Printf(
			"username: %s \npassword: %s \nurl: %s\n",
			crypt.DecryptKey(account.Username, key),
			crypt.DecryptKey(account.Password, key),
			crypt.DecryptKey(account.Url, key),
		)

		// Copy the password to the clipboard
		if runtime.GOOS == "windows" {
			arg = "echo " + crypt.DecryptKey(account.Password, key) + "| clip"
		}

		if runtime.GOOS == "darwin" {
			arg = "echo " + crypt.DecryptKey(account.Password, key) + "| pbcopy"
		}

		cmd := exec.Command("cmd", "/C", arg)
		err := cmd.Start()

		if err != nil {
			fmt.Println(err.Error())
		}

		utils.PrintSuccess("Password copied to clipboard.")
		return
	}

	utils.PrintError("Could not find account with name: " + name)
}

func CreateAccount(newAccount Account) {
	_, found, key := doesAccountExist(newAccount.Name)

	if found == false {
		// Check for password generator key
		if newAccount.Password[:4] == "gen=" {
			fmt.Println("generating password")
			num, err := strconv.Atoi(newAccount.Password[4:])
			if err != nil {
				utils.PrintError("Invalid number ranger please use gen=<valid number>")
				return
			}
			pass := crypt.GenerateKey(num)
			newAccount.Password = pass
		}

		encrypted := Account{
			Name:     crypt.EncryptKey(newAccount.Name, key),
			Username: crypt.EncryptKey(newAccount.Username, key),
			Password: crypt.EncryptKey(newAccount.Password, key),
			Url:      crypt.EncryptKey(newAccount.Url, key),
		}
		strongbox.Accounts = append(strongbox.Accounts, encrypted)
		writeData(strongbox)

		utils.PrintSuccess("account has been created: " + newAccount.Name)
		return
	}

	utils.PrintError("An account with name: " + newAccount.Name + " already exists.")
}

func EditAccount(name string) {
	index, found, key := doesAccountExist(name)

	if found == true {
		account := strongbox.Accounts[index]
		var username string
		var password string
		var url string

		fmt.Println("Edit to update or leave blank to not change.")
		fmt.Printf("username: %s\n", crypt.DecryptKey(account.Username, key))
		fmt.Scan(&username)

		fmt.Printf("password: %s\n", crypt.DecryptKey(account.Password, key))
		fmt.Scan(&password)

		fmt.Printf("url: %s\n", crypt.DecryptKey(account.Url, key))
		fmt.Scan(&url)

		if len(username) != 0 {
			account.Username = crypt.EncryptKey(username, key)
		}

		if len(password) != 0 {
			account.Password = crypt.EncryptKey(password, key)
		}

		if len(url) != 0 {
			account.Url = crypt.EncryptKey(url, key)
		}

		strongbox.Accounts[index] = account
		writeData(strongbox)
		fmt.Printf("%s has been updated.", name)

		return
	}

	utils.PrintError("No account with name: " + name + " exists.")
}

func DeleteAccount(name string) {
	index, found, _ := doesAccountExist(name)

	if found == true {
		strongbox.Accounts = append(strongbox.Accounts[:index], strongbox.Accounts[index+1:]...)
		writeData(strongbox)
		return
	}

	utils.PrintError("No account with name: " + name + " exists.")
}

func doesAccountExist(name string) (index int, found bool, skey string) {
	key := constructKey()
	for index, acc := range strongbox.Accounts {
		if crypt.DecryptKey(acc.Name, key) == name {
			return index, true, key
		}
	}

	return -1, false, key
}
