package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/King12-D/davenated/internal/config"
	"github.com/King12-D/davenated/internal/mailer"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.LoadFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	client := mailer.NewResendClient(cfg.ResendAPIKey)
	if err := mailer.SendEmail(client, cfg.Email); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(os.Stdout, "email sent")
}
