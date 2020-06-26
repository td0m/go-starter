package service

import (
	"github.com/d0minikt/go-starter/server/pkg/domain"
	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", domain.ErrHashing
	}
	return string(hash), nil
}

func matchesHash(plain, hash string) bool {
	bytePlain := []byte(plain)
	byteHash := []byte(hash)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	return err == nil
}