package get

import "errors"

type Link struct {
	Alias string `json:"alias"`
	URL   string `json:"url"`
}

// Repository declares methods required for this service
type Repository interface {
	GetLinks() ([]Link, error)
}

// Service declares functions this service provides
type Service interface {
	GetLinks() []Link
	GetLinkByAlias(string) (*Link, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

// implement your service here
func (s *service) GetLinks() []Link {
	r, _ := s.repo.GetLinks()
	return r
}

func (s *service) GetLinkByAlias(alias string) (*Link, error) {
	links, _ := s.repo.GetLinks()
	for _, l := range links {
		if l.Alias == alias {
			return &l, nil
		}
	}
	return nil, errors.New("link not found")
}
