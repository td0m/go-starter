package add

import "github.com/d0minikt/go-starter/server/pkg/get"

type Link get.Link

// Repository declares methods required for this service
type Repository interface {
	AddLink(Link) error
}

// Service declares functions this service provides
type Service interface {
	AddLink(Link) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

// implement your service here
func (s *service) AddLink(l Link) error {
	return s.repo.AddLink(l)
}
