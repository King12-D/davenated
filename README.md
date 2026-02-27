# Flink Mailer

Branded waitlist email sender for Flink, built in Go with Resend.

It loads `.env`, sends one email per recipient for privacy, and includes rate-limit handling for reliable bulk sends.

## Why this setup

- No recipient address leakage between users.
- Built-in branded Flink HTML template with logo support.
- Retry and pacing logic for Resend `2 requests/second` limits.

## Requirements

- Go `1.25+`
- Resend API key
- Verified sender address on Resend (or allowed sender domain)

## Quick Start

1. Create your env file:
```bash
cp .env.example .env
```
2. Update `.env` values.
3. Run:
```bash
go run ./cmd/flink
```

## Environment Variables

| Variable | Required | Description |
|---|---|---|
| `RESEND_API_KEY` | Yes | Resend API key. |
| `FLINK_EMAIL_TO` | Yes | Comma-separated recipients, e.g. `a@mail.com,b@mail.com`. |
| `FLINK_EMAIL_FROM` | No | Sender format: `email@example.com` or `Name <email@example.com>`. Quote if it contains spaces. Default: `Flink <onboarding@resend.dev>`. |
| `FLINK_EMAIL_SUBJECT` | No | Subject line. Default: `Flink Waitlist Update: We Are Still Building`. |
| `FLINK_EMAIL_LOGO_URL` | No | Public HTTPS logo URL used in the default template. Default: `https://getflink.pro/logo.png`. |
| `FLINK_EMAIL_HTML` | No | Full custom HTML. Leave unset to use the built-in Flink template. |

## Sending Behavior

- Loads `.env` automatically from current directory or parent directories.
- Sends to each recipient individually.
- Adds pacing between sends.
- Retries on rate-limit errors with short backoff.

## Project Structure

- `cmd/flink/main.go` app entrypoint
- `internal/config/config.go` env loading, defaults, and validation
- `internal/mailer/resend.go` Resend client + per-recipient sending

## Common Issues

- `invalid FLINK_EMAIL_FROM "Flink"`: Use a real email format, e.g. `"Flink <info@yourdomain.com>"`.
- `rate limit exceeded`: Keep current logic in place; it already throttles and retries.
- Logo not showing: Ensure `FLINK_EMAIL_LOGO_URL` is a public direct image URL (HTTPS).
