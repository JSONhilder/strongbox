package database

import (
	"fmt"
)

func ListAccounts() {
	if strongbox.Accounts == nil {
		fmt.Println("No accounts in database")
		fmt.Println("Create one with the save command")
		return
	}

	fmt.Println("Stronbox Accounts")
	fmt.Println("-----------------")
	for index, acc := range strongbox.Accounts {
		fmt.Printf("%d: %s\n", index, acc.Name)
	}
}

func GetAccount(name string) {
	for _, acc := range strongbox.Accounts {
		if acc.Name == name {
			if acc.Url == "" {
				acc.Url = "none"
			}
			fmt.Printf("username: %s \npassword: %s \nUrl: %s", acc.Username, acc.Password, acc.Url)
			return // return here to stop extra looping
		}
	}
	fmt.Printf("Could not find account with name: %s", name)
}

func CreateAccount(newAccount Account) {
	_, found := doesAccountExist(newAccount.Name)

	if found == false {
		strongbox.Accounts = append(strongbox.Accounts, newAccount)
		writeData(strongbox)
		return
	}

	fmt.Printf("An account with name: %s already exists.", newAccount.Name)
}

func EditAccount(name string) {
	index, found := doesAccountExist(name)

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
	index, found := doesAccountExist(name)

	if found == true {
		strongbox.Accounts = append(strongbox.Accounts[:index], strongbox.Accounts[index+1:]...)
		writeData(strongbox)
		return
	}

	fmt.Printf("No account with name: %s exists.", name)
}

func doesAccountExist(name string) (index int, found bool) {
	for index, acc := range strongbox.Accounts {
		if acc.Name == name {
			return index, true
		}
	}

	return -1, false
}
