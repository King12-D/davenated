package mailer

import "github.com/resend/resend-go/v3"

type Email struct {
	From    string
	To      []string
	Subject string
	HTML    string
}

func NewResendClient(apiKey string) *resend.Client {
	return resend.NewClient(apiKey)
}

func SendEmail(client *resend.Client, email Email) error {
	_, err := client.Emails.Send(&resend.SendEmailRequest{
		From:    email.From,
		To:      email.To,
		Subject: email.Subject,
		Html:    email.HTML,
	})
	return err
}
