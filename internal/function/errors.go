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

// CustomErr creates a new function error.
func CustomErr(status int, err error) *Error {
	fmt.Println(err) // TODO proper logging
	return &Error{
		error: err,
		code:  status,
	}
}

// Err creates a new function error with the default status text.
func Err(status int, err error) *Error {
	if err != nil { // TODO proper logging
		fmt.Println(err)
	} else {
		fmt.Println(http.StatusText(status))
	}
	return &Error{
		error: errors.New(http.StatusText(status)),
		code:  status,
	}
}
