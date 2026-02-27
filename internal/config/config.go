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

	logoURL := strings.TrimSpace(os.Getenv("FLINK_EMAIL_LOGO_URL"))
	if logoURL == "" {
		logoURL = "https://getflink.pro/logo.png"
	}

	html := strings.TrimSpace(os.Getenv("FLINK_EMAIL_HTML"))
	if html == "" {
		html = defaultHTML(logoURL)
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

func defaultHTML(logoURL string) string {
	return strings.Join([]string{
		"<div style=\"margin:0;padding:0;background:#f3f7f3;font-family:Arial,Helvetica,sans-serif;\">",
		"<div style=\"max-width:620px;margin:0 auto;padding:28px 16px;\">",
		"<div style=\"background:#ffffff;border:1px solid #d7e2d8;border-radius:14px;overflow:hidden;\">",
		"<div style=\"background:linear-gradient(120deg,#0f5d3f,#168b5f);padding:24px 20px;text-align:center;\">",
		"<img src=\"" + logoURL + "\" alt=\"Flink\" style=\"max-height:42px;width:auto;display:inline-block;\"/>",
		"<p style=\"margin:14px 0 0 0;color:#d8f7e8;font-size:13px;\">Waitlist Update</p>",
		"</div>",
		"<div style=\"padding:24px 22px;color:#1d2a22;line-height:1.6;font-size:15px;\">",
		"<p style=\"margin:0 0 14px 0;\">Hi there,</p>",
		"<p style=\"margin:0 0 14px 0;\">Thank you for joining the Flink waitlist at <a href=\"https://getflink.pro\" style=\"color:#0f7a52;text-decoration:none;\">getflink.pro</a>.</p>",
		"<p style=\"margin:0 0 14px 0;\">We are still building Flink carefully so we can launch something useful and reliable for smallholder farmers across Africa.</p>",
		"<p style=\"margin:0 0 14px 0;\">Flink is being built to help farmers decide what to plant, when to plant, and where to sell, using practical information and direct market access.</p>",
		"<p style=\"margin:0 0 10px 0;font-weight:700;\">Current focus areas:</p>",
		"<ul style=\"margin:0 0 14px 20px;padding:0;\">",
		"<li style=\"margin-bottom:8px;\"><strong>Actionable insights:</strong> crop recommendations, seasonal guidance, and local farming tips.</li>",
		"<li style=\"margin-bottom:8px;\"><strong>Market visibility and access:</strong> clearer pricing signals and direct buyer connections.</li>",
		"<li><strong>Low-connectivity reliability:</strong> a lightweight mobile-first experience that works on affordable phones and slower networks.</li>",
		"</ul>",
		"<p style=\"margin:0 0 14px 0;\">Our goal is clear: help farmers improve productivity, reduce waste, and earn more from every harvest.</p>",
		"<p style=\"margin:0 0 14px 0;\"><strong>Launch timeline:</strong> we are preparing to launch in March.</p>",
		"<p style=\"margin:0 0 18px 0;\">We will share another update soon, including early access details.</p>",
		"<p style=\"margin:0 0 14px 0;\"><a href=\"https://getflink.pro\" style=\"display:inline-block;padding:10px 16px;background:#0f7a52;color:#ffffff;text-decoration:none;border-radius:8px;font-weight:600;\">Visit getflink.pro</a></p>",
		"<p style=\"margin:0;\">Thanks for your patience and support.<br/>The Flink Team</p>",
		"</div>",
		"<div style=\"padding:14px 22px;background:#f8fbf9;color:#5f6f64;font-size:12px;border-top:1px solid #e6efe8;\">",
		"<p style=\"margin:0;\">You are receiving this because you joined the Flink waitlist.</p>",
		"</div>",
		"</div>",
		"</div>",
		"</div>",
	}, "\n")
}
