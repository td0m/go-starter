package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/d0minikt/go-starter/server/pkg/domain"
	userHandler "github.com/d0minikt/go-starter/server/pkg/user/delivery"
	userRepo "github.com/d0minikt/go-starter/server/pkg/user/repo"
	userService "github.com/d0minikt/go-starter/server/pkg/user/service"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	cfg := domain.GetConfig()
	db := connectToDB(cfg)
	userRepo := userRepo.New(db)
	userService := userService.New(userRepo)
	userHandler := userHandler.New(userService)
	userRepo.UserExists("email string")

	router := mux.NewRouter()
	router.PathPrefix("/api/users").Handler(userHandler.Handler())
	router.PathPrefix("/auth").Handler(userHandler.Handler())
	router.HandleFunc("/hello", handleHello)
	log.Panic(http.ListenAndServe(":"+cfg.Port, router))
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query()["name"][0]
	fmt.Fprintf(w, "Hello %s", name)
}

func connectToDB(cfg *domain.Config) *sql.DB {
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DB,
	)
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
