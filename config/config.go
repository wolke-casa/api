package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type config struct {
	Port        string   `env:"PORT,required"`
	MaxFileSize int64    `env:"MAX_FILE_SIZE,required"`
	DatabaseUrl string   `env:"DATABASE_URL,required"`
	BotApiKey   string   `env:"BOT_API_KEY,required"`
	KeyLength   int      `env:"KEY_LENGTH,required"`
	IdLength    int      `env:"ID_LENGTH,required"`
	Domains     []string `env:"DOMAINS,required" envSeparator:","`
	Storage     string   `env:"STORAGE,required"`
	Directory   string   `env:"DIRECTORY"`
	AwsBucket   string   `env:"AWS_BUCKET"`
	AwsRegion   string   `env:"AWS_REGION"`
	AwsKey      string   `env:"AWS_KEY"`
	AwsSecret   string   `env:"AWS_SECRET"`
}

var Config config

func Load() error {
	err := godotenv.Load(".env")

	err = env.Parse(&Config)

	return err
}
