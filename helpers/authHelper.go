package helpers

import (
	"log"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panicln(err)
	}
	return string(bytes)
}

func VerifyPassword(actualPassword string, enterdPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(actualPassword), []byte(enterdPassword))

	if err != nil {
		return false, "username or password is incorrect"
	}
	return true, ""
}

func RandNumberGenerator() int {

	rand.New(rand.NewSource(time.Now().UnixNano()))

	min := 100
	max := 1000

	return rand.Intn(max-min+1) + min
}

func PasswordGenerator() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$&"

	b := make([]byte, 8)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
