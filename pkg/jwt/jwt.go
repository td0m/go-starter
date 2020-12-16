package jwt

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type key int

const (
	claimKey key = iota
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

// WithClaims creates the jwt middleware that only accepts traffic that contains a valid jwt token
// errors if invalid token
// puts jwt credentials in context
func (j JWT) WithClaims(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if len(tokenStr) > 0 {
			if !strings.Contains(tokenStr, " ") {
				http.Error(w, "no authorization header", http.StatusUnauthorized)
				return
			}
			parts := strings.Split(tokenStr, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "invalid authorization header. should be in form of: 'Bearer token'", http.StatusUnauthorized)
				return
			}
			tokenStr = parts[1]
		} else {
			tokenStr = r.URL.Query().Get("token")
		}
		if len(tokenStr) == 0 {
			http.Error(w, "authorization token required", http.StatusUnauthorized)
			return
		}
		claims, err := j.getClaims(tokenStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), claimKey, claims)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (j JWT) getClaims(tokenString string) (*Claims, error) {
	claims := Claims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, errors.New("failed parsing token with claims")
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return &claims, nil

}

// FromContext constructs claims from context
func FromContext(ctx context.Context) *Claims {
	return (ctx.Value(claimKey)).(*Claims)
}
