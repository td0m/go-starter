package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/td0m/go-starter/internal/app"
	"github.com/td0m/go-starter/internal/db"
	"github.com/td0m/go-starter/pkg/env"
	"github.com/td0m/go-starter/pkg/migrations"
)

func main() {
	// Postgres
	postgresDB, err := initPostgres(env.PostgresURI)
	check(err)
	check(migrations.ApplyAll(postgresDB, "./sql/schema"))
	fmt.Println("Postgres connected.")

	db := db.New(postgresDB)

	// Routing
	app := app.New(db)
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
