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

// CustomErr creates a new 400 status function error.
func CustomErr(err error) *Error {
	fmt.Println("ERROR: ", err) // TODO proper logging
	return &Error{
		error: err,
		code:  http.StatusBadRequest,
	}
}

// Err creates a new function error with the default status text.
func Err(status int, err error) *Error {
	if err != nil { // TODO proper logging
		fmt.Println("ERROR: ", err)
	} else {
		fmt.Println("ERROR: ", http.StatusText(status))
	}
	return &Error{
		error: errors.New(http.StatusText(status)),
		code:  status,
	}
}
