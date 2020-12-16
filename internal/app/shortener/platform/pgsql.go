package platform

import (
	"github.com/jmoiron/sqlx"
	"github.com/td0m/go-starter/internal/model"
)

// PGSQL struct
type PGSQL struct {
	*sqlx.DB
}

// NewPGSQL creates a new postgres db repository
func NewPGSQL(db *sqlx.DB) *PGSQL {
	return &PGSQL{db}
}

// Create creates a link
func (p PGSQL) Create(id string, url string) error {
	_, err := p.Exec(`
		INSERT INTO links(id,url) VALUES($1,$2)
	`, id, url)
	return err
}

// Get returns link by id
func (p PGSQL) Get(id string) (link model.Link, err error) {
	err = p.DB.Get(&link, `SELECT * FROM links`)
	return
}

// Update updates link by id
func (p PGSQL) Update(id string, url string) error {
	_, err := p.Exec(`UPDATE links SET url=$2 WHERE id=$1`, id, url)
	return err
}
