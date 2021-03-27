package database

import "fmt"

//@TODO all crud operations in here
func GetAccount(name string) {
	fmt.Println("In operations, func GetAccount")
	fmt.Println("Passed name: ", name)
	fmt.Println(database.Mhash)
}

func ListAccounts() {
	if database.Accounts == nil {
		fmt.Println("No accounts in database")
	}
}
