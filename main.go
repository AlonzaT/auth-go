package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := "123456789"
	//fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))
	hashedPass, err := hashPass(pass)
	if err != nil {
		panic(err)
	}

	err = comparePass(pass, hashedPass)
	if err != nil {
		log.Fatalln("Not logged in")
	}

	log.Println("Logged in")
}

//How to hash a password
func hashPass(pass string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error while generating hash: %w", err)
	}
	return bs, nil
}

func comparePass(pass string, hashedPass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(pass))
	if err != nil {
		fmt.Errorf("Invalid password: %w", err)
	}

	return nil
}
