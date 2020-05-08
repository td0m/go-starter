package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/d0minikt/go-starter/server/pkg/add"
	"github.com/d0minikt/go-starter/server/pkg/get"
	_ "github.com/lib/pq"
)

type Link struct {
	Alias string `json:"alias"`
	URL   string `json:"url"`
}

type Repo interface {
	AddLink(add.Link) error
	GetLinks() ([]get.Link, error)
}

type repo struct {
	db *sql.DB
}

func New() Repo {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", "admin", "password", "admin")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return &repo{db}
}

func (r *repo) AddLink(l add.Link) error {
	s, err := r.db.Prepare("INSERT INTO links VALUES($1,$2)")
	if err != nil {
		return err
	}
	_, err = s.Exec(l.Alias, l.URL)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) GetLinks() ([]get.Link, error) {
	rows, err := r.db.Query("select alias, longurl from links")
	if err != nil {
		return nil, err
	}
	links := []get.Link{}
	defer rows.Close()
	for rows.Next() {
		l := get.Link{}
		if err := rows.Scan(&l.Alias, &l.URL); err != nil {
			return nil, err
		}
		links = append(links, l)
	}
	return links, nil
}
