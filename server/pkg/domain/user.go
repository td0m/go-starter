package domain

import "github.com/dgrijalva/jwt-go"

type User struct {
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

type UserRepo interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(email, passwordHash string) error
	RemoveUser(email string) (ok bool)
	UserExists(email string) bool
}

type UserService interface {
	GetUser(email string) (*User, error)
	Register(email string, password string) (*User, error)
	Login(email string, password string) (*User, string, error)
	GetTokenClaims(tokenString string) (jwt.MapClaims, error)
}
