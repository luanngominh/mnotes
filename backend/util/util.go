package util

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"

	"github.com/luanngominh/mnotes/backend/config"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

//GenerateToken generate 6 digits
func GenerateToken() string {
	token := ""
	for i := 0; i < 6; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(10))
		token = token + strconv.Itoa(int(n.Int64()))
	}
	return token
}

//SendVerifyEmail send email to user to verify email
func SendVerifyEmail(name, email, verify string) error {
	from := mail.NewEmail("mnotes", "verify@mnotes.live")
	subject := "Mnotes verify account"
	to := mail.NewEmail(name, email)
	plainTextContent := fmt.Sprintf(
		"Hello %s!<br> We send your verify code: <strong>%s</strong>. Thank you for use our service.",
		name, verify)
	htmlContent := plainTextContent
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(config.Cfg.SendgridAPI)
	_, err := client.Send(message)

	return err
}

//SendWelcomeEmail send email to user to verify email
func SendWelcomeEmail(name, email string) error {
	from := mail.NewEmail("mnotes", "welcome@mnotes.live")
	subject := "Welcome to mnotes"
	to := mail.NewEmail(name, email)
	plainTextContent := fmt.Sprintf(
		"Hello %s!<br> Welcome to mnotes", name)
	htmlContent := plainTextContent
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(config.Cfg.SendgridAPI)
	_, err := client.Send(message)

	return err
}
