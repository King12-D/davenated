package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	"github.com/King12-D/davenated/internal/config"
	"github.com/King12-D/davenated/internal/mailer"
)

func main() {
	_ = loadDotEnv()

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

func loadDotEnv() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	for {
		candidate := filepath.Join(dir, ".env")
		if _, statErr := os.Stat(candidate); statErr == nil {
			return godotenv.Load(candidate)
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return nil
		}
		dir = parent
	}
}
