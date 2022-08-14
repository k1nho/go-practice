package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	// salt rounds
	bytepw, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(bytepw), err
}

func CheckHashPasword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	//password obtained
	password := "secret"

	hash, _ := HashPassword(password)

	fmt.Printf("password is %s \n", password)
	fmt.Printf("hash is %s \n", hash)

	match := CheckHashPasword(password, hash)
	fmt.Printf("Match: %v\n", match)
}
