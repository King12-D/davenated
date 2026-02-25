package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/King12-D/davenated/internal/mailer"
)

type Config struct {
	ResendAPIKey string
	Email        mailer.Email
}

func LoadFromEnv() (Config, error) {
	apiKey := strings.TrimSpace(os.Getenv("RESEND_API_KEY"))
	if apiKey == "" {
		return Config{}, fmt.Errorf("missing RESEND_API_KEY")
	}

	from := strings.TrimSpace(os.Getenv("FLINK_EMAIL_FROM"))
	if from == "" {
		from = "Flink"
	}

	toRaw := strings.TrimSpace(os.Getenv("FLINK_EMAIL_TO"))
	if toRaw == "" {
		return Config{}, fmt.Errorf("missing FLINK_EMAIL_TO")
	}
	to := splitCSV(toRaw)

	subject := strings.TrimSpace(os.Getenv("FLINK_EMAIL_SUBJECT"))
	if subject == "" {
		subject = "Flink waiting list update"
	}

	html := strings.TrimSpace(os.Getenv("FLINK_EMAIL_HTML"))
	if html == "" {
		html = defaultHTML()
	}

	return Config{
		ResendAPIKey: apiKey,
		Email: mailer.Email{
			From:    from,
			To:      to,
			Subject: subject,
			HTML:    html,
		},
	}, nil
}

func splitCSV(value string) []string {
	parts := strings.Split(value, ",")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			continue
		}
		out = append(out, trimmed)
	}
	return out
}

func defaultHTML() string {
	return strings.Join([]string{
		"<p>Hey Dear</p>",
		"<p>Hello,</p>",
		"<p>In January, you joined the Pantero waitlist.</p>",
		"<p>Since then, we have been building deliberately and strategically. Rather than rushing to launch with incomplete features, we chose to focus on designing the system correctly from the ground up.</p>",
		"<p>Pantero is not being built as another content platform. It is being structured as a framework for real skill development and measurable progress.</p>",
		"<p>Over the past few months, our focus has been on three core areas:</p>",
		"<ol>",
		"<li>Structured Learning Paths: Creating guided pathways that move users from foundational knowledge to practical competence, not passive consumption.</li>",
		"<li>Learning to Doing to Earning: Designing a progression system that connects skill acquisition with execution and real opportunities.</li>",
		"<li>Simplicity and Clarity: Ensuring the platform remains focused, outcome-driven, and free from unnecessary complexity.</li>",
		"</ol>",
		"<p>Our goal is simple: when early access begins, users should experience immediate value, not promises of coming soon features.</p>",
	}, "\n")
}
