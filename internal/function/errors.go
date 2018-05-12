package function

import (
	"errors"
	"fmt"
	"net/http"
)

// Error adds a status code to the error type.
type Error struct {
	error
	code int
}

// ClientErr creates a new function error.
func ClientErr(status int, err error) *Error {
	fmt.Println(err) // TODO remove
	return &Error{
		error: err,
		code:  status,
	}
}

// ServerErr creates a new function error and logs it.
func ServerErr(status int, err error) *Error {
	fmt.Println(err) // TODO proper logging
	return &Error{
		error: errors.New(http.StatusText(status)),
		code:  status,
	}
}
