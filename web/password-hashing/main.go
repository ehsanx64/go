package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "secret"
	hash, _ := hashPassword(password)

	fmt.Printf("%-9s: %s\n", "Password", password)
	fmt.Printf("%-9s: %s\n", "Hash", hash)

	match := checkPasswordHash(password, hash)
	fmt.Printf("%-9s: %v\n", "Match", match)
}
