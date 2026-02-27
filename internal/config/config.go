package config

import (
	"fmt"
	"net/mail"
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
		from = "Flink <onboarding@resend.dev>"
	}
	if _, err := mail.ParseAddress(from); err != nil {
		return Config{}, fmt.Errorf("invalid FLINK_EMAIL_FROM %q: use email@example.com or Name <email@example.com>", from)
	}

	toRaw := strings.TrimSpace(os.Getenv("FLINK_EMAIL_TO"))
	if toRaw == "" {
		return Config{}, fmt.Errorf("missing FLINK_EMAIL_TO")
	}
	to := splitCSV(toRaw)
	if len(to) == 0 {
		return Config{}, fmt.Errorf("FLINK_EMAIL_TO has no valid email addresses")
	}
	for _, recipient := range to {
		if _, err := mail.ParseAddress(recipient); err != nil {
			return Config{}, fmt.Errorf("invalid recipient %q in FLINK_EMAIL_TO", recipient)
		}
	}

	subject := strings.TrimSpace(os.Getenv("FLINK_EMAIL_SUBJECT"))
	if subject == "" {
		subject = "Flink Waitlist Update: We Are Still Building"
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
		"<h2>Flink Waitlist Update</h2>",
		"<p>Hi there,</p>",
		"<p>Thank you for joining the Flink waitlist at <a href=\"https://getflink.pro\">getflink.pro</a>.</p>",
		"<p>We are still building Flink carefully so we can launch something useful and reliable for smallholder farmers across Africa.</p>",
		"<p>Flink is being built to help farmers decide what to plant, when to plant, and where to sell, using practical information and direct market access.</p>",
		"<p><strong>Current focus areas:</strong></p>",
		"<ul>",
		"<li><strong>Actionable insights:</strong> crop recommendations, seasonal guidance, and local farming tips.</li>",
		"<li><strong>Market visibility and access:</strong> clearer pricing signals and direct buyer connections.</li>",
		"<li><strong>Low-connectivity reliability:</strong> a lightweight mobile-first experience that works on affordable phones and slower networks.</li>",
		"</ul>",
		"<p>Our goal is clear: help farmers improve productivity, reduce waste, and earn more from every harvest.</p>",
		"<p>We will share another update soon, including early access details.</p>",
		"<p><a href=\"https://getflink.pro\" style=\"display:inline-block;padding:10px 16px;background:#fff;color:#111;text-decoration:none;border-radius:6px;\">Visit getflink.pro</a></p>",
		"<p>Thanks for your patience and support.<br/>The Flink Team</p>",
	}, "\n")
}
