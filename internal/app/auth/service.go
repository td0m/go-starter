package auth

import (
	"context"
	"net/http"

	"github.com/td0m/go-starter/pkg/errors"
	"golang.org/x/oauth2"
)

// Custom errors
var ()

// Service defines a service
type Service struct {
	githubAuth *oauth2.Config
	jwtGen     JWTGenerator
}

// New construcs a new sevice
func New(gh *oauth2.Config, jwtGen JWTGenerator) *Service {
	return &Service{gh, jwtGen}
}

// JWTGenerator generates a jwt
type JWTGenerator func(string) (string, error)

// GithubAuthURL generates a github oauth2 auth url
func (s *Service) GithubAuthURL(state, redirectURL string) string {
	return s.githubAuth.AuthCodeURL(state)
}

// GithubCodeToToken method
func (s *Service) GithubCodeToToken(code string) (string, error) {
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
