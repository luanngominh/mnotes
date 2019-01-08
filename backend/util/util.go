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

//SendEmail send email to user to verify email
func SendEmail(name, email, verify string) error {
	from := mail.NewEmail("mnotes", "verify@mnotes.live")
	subject := "Mnotes verify account"
	to := mail.NewEmail(name, email)
	plainTextContent := fmt.Sprintf(
		"Hello %s!<br> We send your verify code: <strong>%s</strong>. Thank you for use our service.",
		name, verify)
	htmlContent := plainTextContent
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	fmt.Printf(config.Cfg.SendgridAPI)
	client := sendgrid.NewSendClient(config.Cfg.SendgridAPI)
	_, err := client.Send(message)

	return err
}
