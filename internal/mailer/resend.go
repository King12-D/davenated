package mailer

import "fmt"

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
	if len(email.To) == 0 {
		return fmt.Errorf("no recipients provided")
	}

	for _, recipient := range email.To {
		_, err := client.Emails.Send(&resend.SendEmailRequest{
			From:    email.From,
			To:      []string{recipient},
			Subject: email.Subject,
			Html:    email.HTML,
		})
		if err != nil {
			return fmt.Errorf("failed to send to %s: %w", recipient, err)
		}
	}

	return nil
}
