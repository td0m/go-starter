package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/td0m/go-starter/internal/app"
	"github.com/td0m/go-starter/internal/db"
	"github.com/td0m/go-starter/pkg/env"
	"github.com/td0m/go-starter/pkg/jwt"
	"github.com/td0m/go-starter/pkg/migrations"
	"golang.org/x/oauth2"
)

var githubEndpoint = oauth2.Endpoint{
	AuthURL:  "https://github.com/login/oauth/authorize",
	TokenURL: "https://github.com/login/oauth/access_token",
}

func main() {
	// Postgres
	postgresDB, err := initPostgres(env.PostgresURI)
	check(err)
	check(migrations.ApplyAll(postgresDB, "./sql/schema"))
	fmt.Println("Postgres connected.")

	db := db.New(postgresDB)

	githubAuth := &oauth2.Config{
		ClientID:     env.GithubClientID,
		ClientSecret: env.GithubClientSecret,
		Scopes:       []string{"email"},
		RedirectURL:  env.GithubRedirectURL,
		Endpoint:     githubEndpoint,
	}

	go func() {
		time.Sleep(time.Second)

		fmt.Println(githubAuth.AuthCodeURL("state"))

		//tk, err := githubAuth.Exchange(context.Background(), "code")
		//fmt.Println(tk, err)
	}()

	// Routing
	app := app.New(db, githubAuth, jwt.New(env.JWTSecret))
	check(http.ListenAndServe(":"+env.Port, app))
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func initPostgres(uri string) (db *sql.DB, err error) {
	db, err = sql.Open("postgres", uri)
	if err != nil {
		return
	}
	err = db.Ping()
	return
}
