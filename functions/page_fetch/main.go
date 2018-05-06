package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/database"
	"github.com/g-harel/targetblank/internal/database/pages"
	"github.com/g-harel/targetblank/internal/function"
)

func handler(req *function.Request, res *function.Response) {
	page, err := pages.New(database.New()).Get(req.Body)
	switch err.(type) {
	case nil:
		res.Body = page.Page
	case database.ItemNotFoundError:
		res.ClientErr(http.StatusNotFound, err)
	default:
		res.ServerErr(http.StatusInternalServerError, err)
	}
}

func main() {
	lambda.Start(function.New(&function.Config{
		RequireAuth: false,
	}, handler))
}
