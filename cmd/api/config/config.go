package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type config struct {
	Port        string   `env:"PORT"`
	Directory   string   `env:"DIRECTORY"`
	MaxFileSize int64    `env:"MAX_FILE_SIZE"`
	DatabaseUrl string   `env:"DATABASE_URL"`
	BotApiToken string   `env:"BOT_API_TOKEN"`
	KeyLength   int      `env:"KEY_LENGTH"`
	IdLength    int      `env:"ID_LENGTH"`
	Domains     []string `env:"DOMAINS" envSeparator:","`
}

var Config config

func Load() error {
	err := godotenv.Load()

	opts := env.Options{RequiredIfNoDef: true}

	err = env.Parse(&Config, opts)

	return err
}
