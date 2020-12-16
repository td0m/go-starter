package app

import (
	"github.com/gorilla/mux"
	"github.com/td0m/go-starter/internal/app/auth"
	"github.com/td0m/go-starter/internal/app/link"
	"github.com/td0m/go-starter/internal/db"
	"github.com/td0m/go-starter/pkg/jwt"
	"github.com/td0m/go-starter/pkg/middleware"
	"golang.org/x/oauth2"
)

var middlewares = []mux.MiddlewareFunc{
	middleware.ContentTypeJSON,
}

// New creates a new app http handler
func New(db db.Querier, ghauth *oauth2.Config, jwtService *jwt.JWT) *mux.Router {
	r := mux.NewRouter()
	r.Use(middlewares...)
	// api := r.PathPrefix("/api").Subrouter()

	auth.NewHTTP(r.PathPrefix("/auth").Subrouter(), auth.New(ghauth, jwtService.Generate), jwtService.WithClaims)
	link.NewHTTP(r.PathPrefix("/s").Subrouter(), link.New(db))

	return r
}
