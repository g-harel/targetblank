package function

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/g-harel/targetblank/api/internal/token"
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

// CustomErr creates a new 400 status function error.
func CustomErr(err error) *Error {
	fmt.Println("ERROR:", err) // TODO proper logging
	return &Error{
		error: err,
		code:  http.StatusBadRequest,
	}
}

// Err creates a new function error with an encrypted message.
func Err(status int, err error) *Error {
	if err != nil { // TODO proper logging
		fmt.Println("ERROR:", err)
	}
	// Message will be empty if encryption returns an error.
	msg, _ := token.Seal([]byte(err.Error()))
	return &Error{
		error: errors.New(msg),
		code:  status,
	}
}
