package config

import (
	"os"
)

type Config struct {
	DSN            string
	REDIS_ENDPOINT string
}

func Initialize() Config {
	// configFile := os.Getenv("CONFIG_FILE")

	config := Config{
		DSN:            os.Getenv("DSN"),
		REDIS_ENDPOINT: os.Getenv("REDIS_ENDPOINT"),
	}
	return config

}
