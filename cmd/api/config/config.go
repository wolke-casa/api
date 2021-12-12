package config

import "github.com/BurntSushi/toml"

type config struct {
	Port        string
	Directory   string
	MaxFileSize int64
	Database    string
}

var Config config

func Load() error {
	_, err := toml.DecodeFile("config.toml", &Config)

	return err
}
