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

	// MongoDBName specifies the name of the mongo database to use
	MongoDBName string
	// MongoURI specifies the mongo connection string
	MongoURI string
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

	MongoDBName = mustGet("MONGO_INITDB_DATABASE")
	MongoURI = fmt.Sprintf(
		"mongodb://%s:%s@%s/%s?authSource=admin",
		mustGet("MONGO_INITDB_ROOT_USERNAME"),
		mustGet("MONGO_INITDB_ROOT_PASSWORD"),
		get("MONGO_HOST", "localhost"),
		MongoDBName,
	)
}
