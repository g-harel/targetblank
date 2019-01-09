package function

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/g-harel/targetblank/internal/crypto"
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

	// Message payload is encrypted using the application secret.
	// Error message is recoverable for debugging.
	msg, err := crypto.Encrypt([]byte(err.Error()))
	if err != nil {
		msg = http.StatusText(http.StatusInternalServerError)
		log.Printf("ERROR*: %v\n", err)
	}

	return &Error{
		error: errors.New(msg),
		code:  http.StatusInternalServerError,
	}
}
