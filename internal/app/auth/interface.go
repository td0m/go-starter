package auth

import (
	"context"

	"golang.org/x/oauth2"
)

// Service interface
type Service interface {
	GithubAuthURL(string) string
	GithubCodeToToken(string) (string, error)
}

// JWTGenerator generates a jwt
type JWTGenerator func(string) (string, error)

// OAuth authenticates via 3rd party oauth2 services
type OAuth interface {
	AuthCodeURL(string, ...oauth2.AuthCodeOption) string
	Exchange(context.Context, string, ...oauth2.AuthCodeOption) (*oauth2.Token, error)
}
