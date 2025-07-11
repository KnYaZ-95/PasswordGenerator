package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

func getSymbols(specSymbols string) ([]rune, error) {
	result := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	spec := "!@#$%^&*()"

	switch strings.ToLower(specSymbols) {
	case "y":
		result += spec
	case "n":
		// do nothing
	default:
		return []rune{}, errors.New("ValueError")
	}

	runes := []rune(result)
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})

	return runes, nil
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

		symbols, err := getSymbols(chooseSymbols)
		if err != nil {
			fmt.Printf("[!] %s: enter 'y' or 'n' only (not key sensitive)\n\n", err)
			continue
		}
		for i := 0; i < passwordLength; i++ {
			fmt.Printf("%c", symbols[rand.Intn(len(symbols))])
		}

		fmt.Print("\nAnother one? (y/n): ")
		fmt.Scan(&anotherOne)
		if strings.ToLower(anotherOne) == "y" {
			continue
		} else {
			break
		}
	}
}
