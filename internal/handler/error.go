package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

// Common client error messages.
const (
	ErrPageNotFound    = "page not found"
	ErrInvalidEmail    = "invalid email address"
	ErrInvalidPassword = "password is too short"
	ErrGeneric         = "something went wrong"
)

// Error adds a status code to the error type.
type Error struct {
	error
	code int
}

// Code returns the error's http status code.
func (e *Error) Code() int {
	return e.code
}

// TODO add AuthError since it's the most common use of ClientErr.

// ClientErr creates a new 400 status function error.
func ClientErr(format string, a ...interface{}) *Error {
	return &Error{
		error: fmt.Errorf(format, a...),
		code:  http.StatusBadRequest,
	}
}

// InternalErr creates a new function error with an encrypted message.
func InternalErr(format string, a ...interface{}) *Error {
	err := fmt.Errorf(format, a...)

	log.Printf("ERROR: %v\n", err)

	return &Error{
		error: errors.New(ErrGeneric),
		code:  http.StatusInternalServerError,
	}
}
