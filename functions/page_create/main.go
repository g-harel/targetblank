package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/database"
	"github.com/g-harel/targetblank/internal/database/pages"
	"github.com/g-harel/targetblank/internal/function"
)

func handler(req *function.Request, res *function.Response) {
	_, _, err := pages.New(database.New()).Add(req.Body) // TODO send email
	switch err.(type) {
	case nil:
	case database.ValidationError:
		res.ClientErr(http.StatusBadRequest, err)
	default:
		res.ServerErr(http.StatusInternalServerError, err)
	}
}

func main() {
	lambda.Start(function.New(&function.Config{}, handler))
}
