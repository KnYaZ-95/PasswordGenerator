package passwordGenerator

import (
	"errors"
	"math/rand/v2"
)

func GeneratePassword(passwordLength int, specSymbols bool) (string, error) {
	lettersAndDigits := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	specs := "!@#$%^&*()"
	var finalString string

	if passwordLength < 1 {
		return "", errors.New("password length must be greater than zero")
	}

	if specSymbols {
		finalString = lettersAndDigits + specs
	} else {
		finalString = lettersAndDigits
	}

	finalStringLength := len(finalString)
	password := make([]byte, passwordLength)

	for i := 0; i < passwordLength; i++ {
		password[i] = finalString[rand.IntN(finalStringLength)]
	}
	return string(password), nil
}
