package service

import (
	"time"

	"github.com/d0minikt/go-starter/server/pkg/domain"
	"github.com/dgrijalva/jwt-go"
)

func createToken(email string) (string, error) {
	secret := domain.GetConfig().JWTSecret
	claims := jwt.StandardClaims{}
	claims.Subject = email
	claims.ExpiresAt = time.Now().
		Add(time.Hour * 24 * 7).
		Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secret))
	return tokenStr, err
}
