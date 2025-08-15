package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strings"
)

func generatePassword(passwordLength int, specSymbols string) (string, error) {
	lettersAndDigits := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	specs := "!@#$%^&*()"
	var finalString string

	switch strings.ToLower(specSymbols) {
	case "y":
		finalString = lettersAndDigits + specs
	case "n":
		finalString = lettersAndDigits
	default:
		return "", errors.New("ValueError")
	}

	finalStringLength := len(finalString)
	password := make([]byte, passwordLength)

	for i := 0; i < passwordLength; i++ {
		password[i] = finalString[rand.IntN(finalStringLength)]
	}
	return string(password), nil
}

func main() {

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

		fmt.Print("Do you need spec symbols? (y/n): ")
		fmt.Scan(&chooseSymbols)

		password, err := generatePassword(passwordLength, chooseSymbols)
		if err != nil {
			fmt.Printf("[!] %s: enter 'y' or 'n' only (not key sensitive)\n\n", err)
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
