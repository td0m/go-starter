package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/td0m/go-starter/pkg/util/env"
)

func main() {
	_ = initPostgres(env.PostgresURI)

	http.ListenAndServe(":"+env.Port, nil)
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func initPostgres(uri string) (pg *sqlx.DB) {
	i := 0
	var err error
	for i < 5 {
		pg, err = sqlx.Connect("postgres", uri)
		if err == nil {
			log.Println("Postgres connected.")
			return
		}
		i++
		log.Println("Failed to connect to postgres. Trying again...")
		time.Sleep(time.Second)
	}
	log.Panic(err)
	return
}
