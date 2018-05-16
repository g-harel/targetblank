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

	if req.ValidateToken() != nil {
		item, err := pages.New(client).Fetch(addr)
		switch err.(type) {
		case nil:
		case database.ItemNotFoundError:
			return function.Err(http.StatusForbidden, err)
		default:
			return function.Err(http.StatusInternalServerError, err)
		}

		match := hash.Check(req.Body, item.Password)
		if !match {
			return function.Err(http.StatusForbidden, errors.New("password mismatch"))
		}
	}

	return res.SendToken()
}

func main() {
	lambda.Start(function.New(handler))
}
