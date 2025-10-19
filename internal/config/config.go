package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

type ServerConfig struct {
	Port string
}

func Load() *Config {
	// Try to load .env file from current directory first
	envPaths := []string{
		".env",
		"../.env",
	}

	envLoaded := false
	for _, path := range envPaths {
		if err := godotenv.Load(path); err == nil {
			log.Printf("Loaded .env file from: %s", path)
			envLoaded = true
			break
		}
	}

	if !envLoaded {
		log.Println("No .env file found, using environment variables or defaults")
	}

	config := &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DATABASE_HOST", "localhost"),
			Port:     getEnv("DATABASE_PORT", "5432"),
			Database: getEnv("DATABASE_DATABASE", "easy-attend"),
			Username: getEnv("DATABASE_USERNAME", "postgres"),
			Password: getEnv("DATABASE_PASSWORD", "1234"), // Set default password
		},
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
	}

	// Debug logging
	log.Printf("Database config: Host=%s, Port=%s, Database=%s, Username=%s",
		config.Database.Host, config.Database.Port, config.Database.Database, config.Database.Username)

	return config
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
