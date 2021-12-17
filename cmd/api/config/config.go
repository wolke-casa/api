package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type config struct {
	Port        string `env:"PORT"`
	Directory   string `env:"DIRECTORY"`
	MaxFileSize int64  `env:"MAX_FILE_SIZE"`
	DatabaseUrl string `env:"DATABASE_URL"`
	BotApiToken string `env:"BOT_API_TOKEN"`
	KeyLength   int    `env:"KEY_LENGTH"`
	IdLength    int    `env:"ID_LENGTH"`
	Domain      string `env:"DOMAIN"`
}

var Config config

func Load() error {
	// _, err := toml.DecodeFile("config.toml", &Config)
	err := godotenv.Load()

	err = env.Parse(&Config)

	return err
}
