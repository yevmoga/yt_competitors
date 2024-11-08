package configs

import (
	"os"
)

type Config struct {
	ApiKey string
}

func New() (*Config, error) {
	return &Config{
		ApiKey: os.Getenv("YT_API_KEY"),
	}, nil
}
