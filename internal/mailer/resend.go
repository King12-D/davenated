package mailer

import (
	"fmt"
	"strings"
	"time"

	"github.com/resend/resend-go/v3"
)

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

	const sendInterval = 600 * time.Millisecond

	for i, recipient := range email.To {
		if err := sendWithRetry(client, email, recipient); err != nil {
			return fmt.Errorf("failed to send to %s: %w", recipient, err)
		}

		if i < len(email.To)-1 {
			time.Sleep(sendInterval)
		}
	}

	return nil
}

func sendWithRetry(client *resend.Client, email Email, recipient string) error {
	const maxAttempts = 4

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		_, err := client.Emails.Send(&resend.SendEmailRequest{
			From:    email.From,
			To:      []string{recipient},
			Subject: email.Subject,
			Html:    email.HTML,
		})
		if err == nil {
			return nil
		}

		if !isRateLimitError(err) || attempt == maxAttempts {
			return err
		}

		time.Sleep(time.Duration(attempt) * time.Second)
	}

	return nil
}

func isRateLimitError(err error) bool {
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "rate limit exceeded") || strings.Contains(msg, "too many requests")
}
