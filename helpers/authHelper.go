package helpers

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/dipper-lab/asctp-auth/mail"

	"github.com/joho/godotenv"
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

func SendMail(username string, password string, userEmail string) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	emailSenderName := os.Getenv("EMAIL_SENDER_NAME")
	emailSenderAddress := os.Getenv("EMAIL_SENDER_ADDRESS")
	emailSenderPassword := os.Getenv("EMAIL_SENDER_PASSWORD")

	sender := mail.NewGmailSender(emailSenderName, emailSenderAddress, emailSenderPassword)

	subject := "ASCTP Mobile Logins"

	content := fmt.Sprintf(`
	<h1>Here are your logins for the ASCTP Mobile App</h1>
	<p>Username: %s </p>
	<p>Password: %s </p>
	`, username, password)

	to := []string{userEmail}

	err = sender.SendMail(subject, content, to, nil, nil, nil)
	if err != nil {
		fmt.Print(err)
	}
}
