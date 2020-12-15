package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/td0m/go-starter/pkg/util/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// Postgres
	_, err := initPostgres(env.PostgresURI)
	check(err)
	fmt.Println("Postgres connected.")

	// Mongo
	_, err = initMongo(env.MongoDBName, env.MongoURI)
	check(err)
	fmt.Println("Mongo connected.")

	// Routing
	http.ListenAndServe(":"+env.Port, nil)
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func initPostgres(uri string) (*sqlx.DB, error) {
	pg, err := sqlx.Connect("postgres", uri)
	return pg, err
}

func initMongo(dbname, uri string) (db *mongo.Database, err error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return
	}

	pingCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	err = client.Ping(pingCtx, readpref.Primary())
	if err != nil {
		return
	}

	return client.Database(dbname), err
}
