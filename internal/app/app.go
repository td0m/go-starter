package app

import (
	"github.com/gorilla/mux"
	"github.com/td0m/go-starter/internal/app/link"
	"github.com/td0m/go-starter/internal/db"
	"github.com/td0m/go-starter/pkg/middleware"
)

var middlewares = []mux.MiddlewareFunc{
	middleware.ContentTypeJSON,
}

// New creates a new app http handler
func New(db db.Querier) *mux.Router {
	r := mux.NewRouter()
	r.Use(middlewares...)
	// api := r.PathPrefix("/api").Subrouter()

	link.NewHTTP(r.PathPrefix("/s").Subrouter(), link.New(db))

	return r
}
