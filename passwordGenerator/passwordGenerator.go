package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	lowercase    = "abcdefghijklmnopqrstuvwxyz"
	uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers      = "0123456789"
	specialChars = "!@#$%^&*()-_=+[]{}|;:,.<>?/`~"
)

func generatePassword(length int) (string, error) {
	charSet := lowercase + uppercase + numbers + specialChars

	password := make([]byte, length)
	for i := range password {
		randomChar, err := getRandomChar(charSet)
		if err != nil {
			return "", err
		}
		password[i] = randomChar
	}

	return string(password), nil
}

func getRandomChar(charSet string) (byte, error) {
	max := big.NewInt(int64(len(charSet)))
	randomInt, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return charSet[randomInt.Int64()], nil
}

func main() {
	length := 12

	password, err := generatePassword(length)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Generated Password:", password)
}
