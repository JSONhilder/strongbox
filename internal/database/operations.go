package database

import (
	"fmt"

	"github.com/JSONhilder/strongbox/internal/crypt"
)

func ListAccounts() {
	key := constructKey()
	if strongbox.Accounts == nil {
		fmt.Println("No accounts in database")
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
	index, found, key := doesAccountExist(name)

	if found == true {
		account := strongbox.Accounts[index]
		fmt.Printf(
			"username: %s \npassword: %s \nurl: %s",
			crypt.DecryptKey(account.Username, key),
			crypt.DecryptKey(account.Password, key),
			crypt.DecryptKey(account.Url, key),
		)
		return
	}
	fmt.Printf("Could not find account with name: %s", name)
}

func CreateAccount(newAccount Account) {
	_, found, key := doesAccountExist(newAccount.Name)

	if found == false {
		encrypted := Account{
			Name:     crypt.EncryptKey(newAccount.Name, key),
			Username: crypt.EncryptKey(newAccount.Username, key),
			Password: crypt.EncryptKey(newAccount.Password, key),
			Url:      crypt.EncryptKey(newAccount.Url, key),
		}
		strongbox.Accounts = append(strongbox.Accounts, encrypted)
		writeData(strongbox)
		return
	}

	fmt.Printf("An account with name: %s already exists.", newAccount.Name)
}

func EditAccount(name string) {
	index, found, _ := doesAccountExist(name)

	if found == true {
		account := strongbox.Accounts[index]
		var username string
		var password string
		var url string

		fmt.Println("Edit to update or leave blank to not change.")
		fmt.Printf("username: %s\n", account.Username)
		fmt.Scanln(&username)

		fmt.Printf("password: %s\n", account.Password)
		fmt.Scanln(&password)

		fmt.Printf("url: %s\n", account.Url)
		fmt.Scanln(&url)

		if len(username) != 0 {
			account.Username = username
		}

		if len(password) != 0 {
			account.Password = password
		}

		if len(url) != 0 {
			account.Url = url
		}

		strongbox.Accounts[index] = account
		writeData(strongbox)

		return
	}

	fmt.Printf("No account with name: %s exists.", name)

}

func DeleteAccount(name string) {
	index, found, _ := doesAccountExist(name)

	if found == true {
		strongbox.Accounts = append(strongbox.Accounts[:index], strongbox.Accounts[index+1:]...)
		writeData(strongbox)
		return
	}

	fmt.Printf("No account with name: %s exists.", name)
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
