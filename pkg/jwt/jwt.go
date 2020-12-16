package jwt

import (
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

// Claims holds all custom claims
type Claims struct {
	jwt.StandardClaims
}

// JWT struct
type JWT struct {
	method jwt.SigningMethod
	secret string
}

// New creates a new JWT
func New(secret string) *JWT {
	return &JWT{
		jwt.SigningMethodHS256,
		secret,
	}
}

// Generate generates a new jwt token
func (j JWT) Generate(id string) (string, error) {
	claims := Claims{
		StandardClaims: jwt.StandardClaims{Subject: id, ExpiresAt: 0},
	}
	token := jwt.NewWithClaims(j.method, claims)
	return token.SignedString([]byte(j.secret))
}

// Middleware creates the jwt middleware
func (j JWT) Middleware() func(http.Handler) http.Handler {
	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(t *jwt.Token) (interface{}, error) {
			return []byte(j.secret), nil
		},
		SigningMethod: j.method,
	}).Handler
}
