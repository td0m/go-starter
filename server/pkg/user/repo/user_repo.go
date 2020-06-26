package repo

import (
	"database/sql"

	"github.com/d0minikt/go-starter/server/pkg/domain"
)

type repo struct {
	Conn *sql.DB
}

func New(conn *sql.DB) domain.UserRepo {
	return &repo{conn}
}

func (r *repo) GetUserByEmail(email string) (*domain.User, error) {
	u := domain.User{}
	err := r.Conn.
		QueryRow("SELECT email,password_hash FROM users WHERE email=$1", email).
		Scan(&u.Email, &u.PasswordHash)
	if err != nil {
		return nil, domain.ErrNotFound
	}
	return &u, nil
}

func (r *repo) CreateUser(email, passwordHash string) error {
	q, err := r.Conn.Prepare("INSERT INTO users(email,password_hash) VALUES($1,$2)")
	if err != nil {
		return domain.ErrInternal
	}
	_, err = q.Exec(email, passwordHash)
	if err != nil {
		return domain.ErrInsertConflict
	}
	return nil
}

func (r *repo) UserExists(email string) bool {
	return true
}

func (r *repo) RemoveUser(email string) bool {
	return false
}
