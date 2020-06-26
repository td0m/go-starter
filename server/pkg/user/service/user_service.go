package service

import (
	"errors"
	"fmt"

	"github.com/d0minikt/go-starter/server/pkg/domain"
	"github.com/dgrijalva/jwt-go"
)

type service struct {
	repo domain.UserRepo
}

func New(repo domain.UserRepo) domain.UserService {
	return &service{repo}
}

func (s *service) GetUser(email string) (*domain.User, error) {
	return s.repo.GetUserByEmail(email)
}

func (s *service) Register(email, password string) (*domain.User, error) {
	hash, err := hashAndSalt(password)
	if err != nil {
		return nil, err
	}
	s.repo.CreateUser(email, hash)
	return nil, nil
}

func (s *service) Login(email string, password string) (*domain.User, string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, "", err
	}
	if !matchesHash(password, user.PasswordHash) {
		return nil, "", domain.ErrCredentials
	}
	token, err := createToken(email)
	if err != nil {
		return nil, "", err
	}
	return user, token, nil
}

func (s *service) GetTokenClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(domain.GetConfig().JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	/*
		if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
			return nil, err
		}
	*/
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("Invalid token")
	}
	return claims, nil
}
