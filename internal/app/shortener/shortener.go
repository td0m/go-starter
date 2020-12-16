package shortener

import "github.com/td0m/go-starter/internal/model"

// Get method
func (s Shortener) Get(id string) (model.Link, error) {
	return s.db.Get(id)
}

// Create method
func (s Shortener) Create(id, url string) (model.Link, error) {
	link := model.Link{
		ID:  id,
		URL: url,
	}
	return link, s.db.Create(link.ID, link.URL)
}
