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

	// GithubClientID is the github client app id
	GithubClientID string

	// GithubClientSecret is the github client app secret
	GithubClientSecret string

	// GithubRedirectURL specifies the url to redirect to after obtaining oauth2 auth code
	GithubRedirectURL string

	// JWTSecret is the jwt signing secret
	JWTSecret string
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

	GithubClientID = mustGet("GITHUB_CLIENT_ID")
	GithubClientSecret = mustGet("GITHUB_CLIENT_SECRET")
	GithubRedirectURL = mustGet("GITHUB_REDIRECT_URL")

	JWTSecret = mustGet("JWT_SECRET")
}
