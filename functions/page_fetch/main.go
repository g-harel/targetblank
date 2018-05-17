package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/database"
	"github.com/g-harel/targetblank/internal/database/pages"
	"github.com/g-harel/targetblank/internal/function"
)

var client = database.New()

func handler(req *function.Request, res *function.Response) *function.Error {
	addr, err := req.Param("address")
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}

	item, err := pages.New(client).Fetch(addr)
	switch err.(type) {
	case nil:
	case database.ItemNotFoundError:
		return function.Err(http.StatusNotFound, err)
	default:
		return function.Err(http.StatusInternalServerError, err)
	}

	if !item.Published {
		_, funcErr := req.ValidateToken(addr)
		if err != nil {
			return funcErr
		}
	}

	res.Body = item.Page
	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
