package auth

import (
	"context"
	"net/http"

	"github.com/td0m/go-starter/pkg/errors"
)

// Custom errors
var ()

// Auth defines a service
type Auth struct {
	githubAuth OAuth
	jwtGen     JWTGenerator
}

// New construcs a new sevice
func New(gh OAuth, jwtGen JWTGenerator) *Auth {
	return &Auth{gh, jwtGen}
}

// GithubAuthURL generates a github oauth2 auth url
func (s *Auth) GithubAuthURL(state string) string {
	return s.githubAuth.AuthCodeURL(state)
}

// GithubCodeToToken method
func (s *Auth) GithubCodeToToken(code string) (string, error) {
	oauthResp, err := s.githubAuth.Exchange(context.Background(), code)
	if err != nil || !oauthResp.Valid() {
		return "", errors.New(http.StatusUnauthorized, "authentication failed. provided auth code or server configuration is invalid.")
	}
	data, err := getGithubData(oauthResp.AccessToken)
	if err != nil {
		return "", errors.New(http.StatusUnauthorized, "auth token obtained but failed to request user data.")
	}
	return s.jwtGen(data.Login)
}
