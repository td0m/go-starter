package shortener

import (
	"github.com/jmoiron/sqlx"
	"github.com/td0m/go-starter/internal/app/shortener/platform"
	"github.com/td0m/go-starter/internal/model"
)

// Service defines methods in Shortener
type Service interface {
	Get(string) (model.Link, error)
	Create(string, string) (model.Link, error)
}

// Shortener struct
type Shortener struct {
	db DB
}

// New constructs a Shortener
func New(db DB) *Shortener {
	return &Shortener{db}
}

// Init creates a new service using default platforms
func Init(pg *sqlx.DB) *Shortener {
	return New(platform.NewPGSQL(pg))
}

// DB represents database interface methods
type DB interface {
	Create(id, url string) error
	Get(id string) (model.Link, error)
	Update(id, url string) error
}
