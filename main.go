package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"log"
	"time"

	//"github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserClaims struct {
	jwt.StandardClaims
	SessionID int64
}

func (u *UserClaims) Valid() error {
	if !u.StandardClaims.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("token has expired")
	}

	if u.SessionID == 0 {
		return fmt.Errorf("Invalid session ID")
	}

	return nil
}

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

//For the sha512 function needs a key of 64 bytes
var key string = "thiswillblowyourmindtimeaftertimethiswiwetrhjlkdsfgbnmrmindtimeaf"

//this function is using hmac to create a bearer token
//a token or signature to allow easy authorization
func signMsg(msg []byte) ([]byte, error) {
	h := hmac.New(sha512.New, []byte(key))

	_, err := h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("Error while hashing message: %w", err)
	}

	signature := h.Sum(nil)
	return signature, nil
}

//For checking signature
func checkSig(msg, sig []byte) (bool, error) {
	newSig, err := signMsg(msg)
	if err != nil {
		return false, fmt.Errorf("Error in the checking signature: %w", err)
	}

	same := hmac.Equal(newSig, sig)

	return same, nil
}

//Creating Json web tokens
func createToken(c *UserClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)

	signedToken, err := t.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("Error in createToken: %w", err)
	}

	return signedToken, nil

}
