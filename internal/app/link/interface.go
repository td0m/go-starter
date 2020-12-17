package link

import (
	"context"

	"github.com/td0m/go-starter/internal/db"
)

type Service interface {
	Get(string) (*db.Link, error)
	Create(string, string) (*db.Link, error)
}

type DB interface {
	GetLink(context.Context, string) (db.Link, error)
	CreateLink(context.Context, db.CreateLinkParams) (db.Link, error)
}
