package domain

import "errors"

var (
	ErrInternal       = errors.New("Internal Server Error")
	ErrInsertConflict = errors.New("Item already exists")
	ErrNotFound       = errors.New("Item not found")
	ErrHashing        = errors.New("Hashing failed")
	ErrCredentials    = errors.New("Invalid credentials")
)
