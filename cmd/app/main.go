package main

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"time"

	_ "github.com/lib/pq"
	"github.com/td0m/go-starter/internal/app"
	"github.com/td0m/go-starter/internal/db"
	"github.com/td0m/go-starter/pkg/util/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// Postgres
	postgresDB, err := initPostgres(env.PostgresURI)
	check(err)
	check(migrate(postgresDB, "./sql/schema"))
	fmt.Println("Postgres connected.")

	db := db.New(postgresDB)

	// Mongo
	mongo, err := initMongo(env.MongoDBName, env.MongoURI)
	check(err)
	fmt.Println("Mongo connected.", mongo)

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

func migrate(db *sql.DB, dir string) (err error) {
	migrations, err := getMigrations(dir)
	if err != nil {
		return
	}
	for _, name := range migrations {
		err = apply(db, path.Join(dir, name))
		if err != nil {
			return
		}
	}
	return
}

// apply applies a migration in a given file to the database
func apply(db *sql.DB, path string) error {
	migration, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file at \"%s\"", path)
	}

	if _, err := db.Exec(string(migration)); err != nil {
		return fmt.Errorf("error applying migration \"%s\": %s", path, err)
	}
	return nil
}

func getMigrations(dir string) (out []string, err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	for _, f := range files {
		out = append(out, f.Name())
	}
	return
}
