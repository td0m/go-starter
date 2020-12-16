package env

import (
	"fmt"

	"github.com/joho/godotenv"
)

var (
	// Port is the http port the server will listen on
	Port string

	// PostgresURI specifies the postgres connection string
	PostgresURI string
)

func init() {
	godotenv.Load()

	Port = get("PORT", "8080")

	PostgresURI = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		get("POSTGRES_HOST", "localhost"),
		mustGet("POSTGRES_USER"),
		mustGet("POSTGRES_PASSWORD"),
		mustGet("POSTGRES_DB"),
	)
}
