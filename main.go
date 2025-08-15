package main

import (
	"PasswordGenerator/passwordGenerator"
	"fmt"
	"strings"
)

func main() {

	var password string
	var passwordLength int
	var chooseSymbols string
	var anotherOne string

	for {
		fmt.Print("Enter password length: ")
		_, err := fmt.Scan(&passwordLength)
		if err != nil || passwordLength < 1 {
			fmt.Print("[!] Please enter a valid positive number\n\n")
			continue
		}

		fmt.Print("Do you need special symbols? (y/n): ")
		fmt.Scan(&chooseSymbols)

		switch strings.ToLower(chooseSymbols) {
		case "y":
			password, _ = passwordGenerator.GeneratePassword(passwordLength, true)
		case "n":
			password, _ = passwordGenerator.GeneratePassword(passwordLength, false)
		default:
			fmt.Println("[!] Enter 'y' or 'n' only (not key sensitive)")
			continue
		}

		fmt.Printf("%v\nAnother one? (y/n): ", password)
		fmt.Scan(&anotherOne)
		if strings.ToLower(anotherOne) == "y" {
			continue
		} else {
			break
		}
	}
}
