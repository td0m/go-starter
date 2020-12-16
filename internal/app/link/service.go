package link

import (
	"context"
	"net/http"

	"github.com/td0m/go-starter/internal/db"
	"github.com/td0m/go-starter/pkg/errors"
)

// Custom errors
var (
	ErrNotFound      = errors.New(http.StatusNotFound, "link not found")
	ErrAlreadyExists = errors.New(http.StatusConflict, "link already exists")
)

// Service defines a service
type Service struct {
	db db.Querier
}

// New construcs a new sevice
func New(db db.Querier) *Service {
	return &Service{db}
}

// Get method
func (s Service) Get(id string) (*db.Link, error) {
	link, err := s.db.GetLink(context.Background(), id)
	if err != nil {
		return &link, ErrNotFound
	}
	return &link, nil
}

// Create method
func (s Service) Create(id, url string) (*db.Link, error) {
	l, err := s.db.CreateLink(context.Background(), db.CreateLinkParams{
		ID:  id,
		Url: url,
	})
	if err != nil {
		return nil, ErrAlreadyExists
	}
	return &l, nil
}
