package main

import (
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/database"
	"github.com/g-harel/targetblank/internal/database/pages"
	"github.com/g-harel/targetblank/internal/function"
	"github.com/g-harel/targetblank/internal/hash"
)

var client = database.New()

func handler(req *function.Request, res *function.Response) *function.Error {
	addr, err := req.Param("address")
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}

	if req.Body != "" {
		item, err := pages.New(client).Fetch(addr)
		switch err.(type) {
		case nil:
		case database.ItemNotFoundError:
			return function.Err(http.StatusForbidden, err)
		default:
			return function.Err(http.StatusInternalServerError, err)
		}

		if !hash.Check(req.Body, item.Password) {
			return function.Err(http.StatusForbidden, errors.New("password mismatch"))
		}
	} else {
		restricted, funcErr := req.ValidateToken(addr)
		if funcErr != nil {
			return funcErr
		}
		if restricted {
			return function.CustomErr(errors.New("cannot refresh restricted token"))
		}
	}

	token, funcErr := function.MakeToken(false, addr)
	if err != nil {
		return funcErr
	}

	req.Body = token
	req.Headers["Content-Type"] = "text/plain"

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
