package utils

import "fmt"

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"

func PrintSuccess(msg string) {
	fmt.Println(Green + msg + Reset)
}

func PrintError(msg string) {
	fmt.Println(Red + msg + Reset)
}

func PrintWarn(msg string) {
	fmt.Println(Yellow + msg + Reset)
}
