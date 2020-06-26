package domain

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type PostgresConfig struct {
	User, Password, DB, Host string
}

type Config struct {
	Production bool
	Port       string
	JWTSecret  string
	Postgres   PostgresConfig
}

var config Config

func GetConfig() *Config {
	if len(config.Port) == 0 {
		config = initConfig()
	}
	return &config
}

func initConfig() Config {
	godotenv.Load()
	prod := has("PRODUCTION")
	port := get("PORT", "8080")
	dbHost := "localhost"

	if prod {
		dbHost = "db"
	}

	return Config{
		Production: prod,
		Port:       port,
		JWTSecret:  mustGet("JWT_SECRET"),
		Postgres: PostgresConfig{
			User:     mustGet("POSTGRES_USER"),
			Password: mustGet("POSTGRES_PASSWORD"),
			DB:       mustGet("POSTGRES_DB"),
			Host:     dbHost,
		},
	}
}

func mustGet(key string) string {
	v := os.Getenv(key)
	if len(v) == 0 {
		log.Fatal("Failed to get env variable: " + key)
	}
	return v
}

func get(key, defaultValue string) string {
	v := os.Getenv(key)
	if len(v) == 0 {
		return defaultValue
	}
	return v
}

func has(key string) bool {
	_, ok := os.LookupEnv(key)
	return ok
}
