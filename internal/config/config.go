package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	ConnectionString string
}

type JWTConfig struct {
	Secret string
	Expiry time.Duration
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	port, err := getRequiredEnv("PORT")
	if err != nil {
		return nil, err
	}

	connectionString, err := getRequiredEnv("DATABASE_CONNECTION_STRING")

	fmt.Println("DATABASE_CONNECTION_STRING =", os.Getenv("DATABASE_CONNECTION_STRING"))
	
	if err != nil {
		return nil, err
	}

	jwtSecret, err := getRequiredEnv("JWT_SECRET")
	if err != nil {
		return nil, err
	}

	jwtExpiryString, err := getRequiredEnv("JWT_EXPIRY")
	if err != nil {
		return nil, err
	}

	jwtExpiry, err := time.ParseDuration(jwtExpiryString)
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_EXPIRY: %w", err)
	}

	return &Config{
		Server: ServerConfig{
			Port: port,
		},
		Database: DatabaseConfig{
			ConnectionString: connectionString,
		},
		JWT: JWTConfig{
			Secret: jwtSecret,
			Expiry: jwtExpiry,
		},
	}, nil
}

func getRequiredEnv(key string) (string, error) {
	value := os.Getenv(key)

	if value == "" {
		return "", fmt.Errorf("missing required environment variable: %s", key)
	}

	return value, nil
}
