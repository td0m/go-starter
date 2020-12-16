package errors

import (
	"errors"
	"fmt"
	"net/http"
)

// Error represents the type of a http error
type Error struct {
	Code int
	Err  error
}

func (e Error) Error() string {
	return fmt.Sprintf("%s", e.Err.Error())
}

// Write writes a response containing information about the error
// If it's an unexpected internal error, it will log it internally instead
func (e Error) Write(w http.ResponseWriter) {
	http.Error(w, e.Error(), e.Code)
}

// New creates a new Error
func New(code int, s string) *Error {
	return &Error{
		Err:  errors.New(s),
		Code: code,
	}
}

// WithCode creates a custom error with code
func WithCode(code int, err error) *Error {
	return New(code, err.Error())
}

func internal(err error) *Error {
	return &Error{Code: http.StatusInternalServerError, Err: err}
}

// From constructs a new error
func From(anyErr error) *Error {
	switch err := anyErr.(type) {
	case *Error:
		return err
	default:
		return internal(err)
	}
}

// Write utility function
func Write(w http.ResponseWriter, err error) {
	From(err).Write(w)
}
