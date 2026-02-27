# Flink Mailer

Small Go CLI for sending a Flink waitlist update email using Resend.

## Requirements

- Go 1.25+
- A Resend API key

## Setup

1. Copy `.env.example` to `.env` and fill in the values.
2. The app auto-loads `.env` at startup, so you can run directly.

Required:

- `RESEND_API_KEY`
- `FLINK_EMAIL_TO` (comma-separated list). Example: `person1@example.com,person2@example.com`

Optional:

- `FLINK_EMAIL_FROM` (default: `Flink <onboarding@resend.dev>`, must be `email@example.com` or `Name <email@example.com>`; quote it in `.env` if it contains spaces)
- `FLINK_EMAIL_SUBJECT` (default: `Flink waiting list update`)
- `FLINK_EMAIL_HTML` (default: built-in HTML message)

## Run

```bash
go run ./cmd/flink
```

## Project Layout

- `cmd/flink/main.go` entrypoint
- `internal/config` env parsing and defaults
- `internal/mailer` email sending via Resend
