package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type config struct {
	Port        string
	Directory   string
	MaxFileSize int64
	DatabaseUrl string
	BotApiToken string
	KeyLength   int
	Domain      string
}

var Config config

func Load() error {
	// _, err := toml.DecodeFile("config.toml", &Config)
	err := godotenv.Load()

	conf := &Config

	conf.Port = os.Getenv("PORT")
	conf.Directory = os.Getenv("DIRECTORY")

	parsedMaxFileSize, err := strconv.ParseInt(os.Getenv("MAX_FILE_SIZE"), 10, 64)
	conf.MaxFileSize = parsedMaxFileSize

	conf.DatabaseUrl = os.Getenv("DATABASE_URL")
	conf.BotApiToken = os.Getenv("BOT_API_TOKEN")

	parsedKeyLength, err := strconv.Atoi(os.Getenv("KEY_LENGTH"))
	conf.KeyLength = parsedKeyLength

	conf.Domain = os.Getenv("DOMAIN")

	return err
}
