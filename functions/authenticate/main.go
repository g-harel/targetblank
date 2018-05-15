package main

import (
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/database"
	"github.com/g-harel/targetblank/internal/database/pages"
	"github.com/g-harel/targetblank/internal/function"
	"github.com/g-harel/targetblank/internal/hash"
	"github.com/g-harel/targetblank/internal/token"
)

var client = database.New()

func handler(req *function.Request, res *function.Response) *function.Error {
	addr, err := req.Param("address")
	if err != nil {
		return function.ServerErr(http.StatusInternalServerError, err)
	}

	forbiddenErr := errors.New(http.StatusText(http.StatusForbidden))

	item, err := pages.New(client).Fetch(addr)
	switch err.(type) {
	case nil:
	case database.ItemNotFoundError:
		return function.ClientErr(http.StatusForbidden, forbiddenErr)
	default:
		return function.ServerErr(http.StatusInternalServerError, err)
	}

	match := hash.Check(req.Body, item.Password)
	if !match {
		return function.ClientErr(http.StatusForbidden, forbiddenErr)
	}

	res.Headers["Content-Type"] = "text/plain"
	res.Body = token.Generate(addr)

	return nil

}

func main() {
	lambda.Start(function.New(handler))
}
