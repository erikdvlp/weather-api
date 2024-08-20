package config

import (
	"os"
)

type Config struct {
	Port string
}

func LoadConfig() *Config {
	var config Config
	config.Port = getEnv("PORT", "8080")
	return &config
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
