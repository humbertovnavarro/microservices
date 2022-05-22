package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	Port string
}

type DatabaseConfig struct {
	Port     string
	User     string
	Name     string
	Password string
	Host     string
}

type Config struct {
	Api      ApiConfig
	Database DatabaseConfig
}

func AssertGetenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Panicf("Could not get required environment variable %s", key)
	}
	return value
}

func loadApiConfig() *ApiConfig {
	Port := AssertGetenv("API_PORT")
	config := &ApiConfig{
		Port,
	}
	return config
}

func loadDatabaseConfig() *DatabaseConfig {
	Port := AssertGetenv("DB_PORT")
	User := AssertGetenv("DB_USER")
	Name := AssertGetenv("DB_NAME")
	Password := AssertGetenv("DB_PASS")
	Host := AssertGetenv("DB_HOST")
	databaseConfig := &DatabaseConfig{
		Port,
		User,
		Name,
		Password,
		Host,
	}
	return databaseConfig
}

func LoadConfig() *Config {
	godotenv.Load()
	return &Config{
		Api:      *loadApiConfig(),
		Database: *loadDatabaseConfig(),
	}
}
