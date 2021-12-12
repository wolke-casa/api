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
	Database    string
}

var Config config

func Load() error {
	// _, err := toml.DecodeFile("config.toml", &Config)
	err := godotenv.Load()

	// TODO: This is a hacky fix but itll work for now
	conf := &Config

	conf.Port = os.Getenv("PORT")
	conf.Directory = os.Getenv("DIRECTORY")

	parsedMaxFileSize, err := strconv.ParseInt(os.Getenv("MAXFILESIZE"), 10, 64)
	conf.MaxFileSize = parsedMaxFileSize

	conf.Database = os.Getenv("DATABASE")

	return err
}
