package helpers

import (
	"fmt"
	"log"
	"os"

	"github.com/dipper-lab/asctp-esntls/mail"
	"github.com/joho/godotenv"
)

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

func SendVerificationEMail(name string, token string, email string) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	emailSenderName := os.Getenv("EMAIL_SENDER_NAME")
	emailSenderAddress := os.Getenv("EMAIL_SENDER_ADDRESS")
	emailSenderPassword := os.Getenv("EMAIL_SENDER_PASSWORD")

	sender := mail.NewGmailSender(emailSenderName, emailSenderAddress, emailSenderPassword)

	subject := "AgroTrace Email Verification"

	content := fmt.Sprintf(`
	<h2 style="color: #007bff;">Verify Your Account</h2>

        <p>Hello %s,</p>

        <p>Welcome to AgroTrace. To complete the registration process, please click the link below to verify your email address:</p>

        <p><a href="https://asctp-auth.onrender.com/account/verifiy/%s/%s" style="display: inline-block; padding: 10px 20px; background-color: #007bff; color: #fff; text-decoration: none; border-radius: 5px;">Verify Now</a></p>

        <p>If you did not sign up for an account, you can ignore this email.</p>

        <p>Thank you,<br>DIPPER Lab</p>

	`, name, email, token)

	to := []string{email}

	err = sender.SendMail(subject, content, to, nil, nil, nil)
	if err != nil {
		fmt.Print(err)
	}
}
