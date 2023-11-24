package config

import (
	"os"
)

type Config struct {
	DSN            string
	REDIS_ENDPOINT string
	APP_PORT       string
}

func Initialize() Config {
	// configFile := os.Getenv("CONFIG_FILE")

	config := Config{
		DSN:            os.Getenv("DSN"),
		REDIS_ENDPOINT: os.Getenv("REDIS_ENDPOINT"),
		APP_PORT:       os.Getenv("APP_PORT"),
	}
	return config

}
