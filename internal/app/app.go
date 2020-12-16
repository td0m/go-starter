package app

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/td0m/go-starter/internal/app/shortener"
	shortenerD "github.com/td0m/go-starter/internal/app/shortener/delivery"
	"go.mongodb.org/mongo-driver/mongo"
)

// New creates a new app http handler
func New(postgres *sqlx.DB, mongo *mongo.Database) *mux.Router {

	r := mux.NewRouter()
	// api := r.PathPrefix("/api").Subrouter()

	shortenerD.New(r.PathPrefix("/s").Subrouter(), shortener.Init(postgres))

	return r
}
