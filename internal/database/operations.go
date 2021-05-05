package database

import (
	"fmt"
)

//@TODO all crud operations in here
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

}

func DeleteAccount(name string) {
	index, found := doesAccountExist(name)

	if found == true {
		strongbox.Accounts = append(strongbox.Accounts[:index], strongbox.Accounts[index+1:]...)
		writeData(strongbox)
		return
	}
}

func doesAccountExist(name string) (index int, found bool) {
	for index, acc := range strongbox.Accounts {
		if acc.Name == name {
			return index, true
		}
	}
	fmt.Printf("No account with name: %s exists.", name)
	return -1, false
}
