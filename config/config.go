package config

import (
	"log"
	"os"
)

type Config struct {
	DBHost, DBPort, DBUser, DBPassword, DBName string
}

var AppConfig Config

// Load environment variables
func Load() {
	AppConfig.DBHost = getEnv("DB_HOST", "localhost")
	AppConfig.DBPort = getEnv("DB_PORT", "5433")
	AppConfig.DBUser = getEnv("DB_USER", "admin")
	AppConfig.DBPassword = getEnv("DB_PASSWORD", "admin")
	AppConfig.DBName = getEnv("DB_NAME", "blackenshovel")

	log.Println("Configuration loaded")
}

// getEnv returns default if variable not set
func getEnv(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}
