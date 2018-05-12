package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/database"
	"github.com/g-harel/targetblank/internal/database/pages"
	"github.com/g-harel/targetblank/internal/function"
)

func handler(req *function.Request, res *function.Response) {
	pass := database.RandString(16)
	_, err := pages.New(database.New()).Create(req.Body, pass)
	switch err.(type) {
	case nil:
	case database.ValidationError:
		res.ClientErr(http.StatusBadRequest, err)
	default:
		res.ServerErr(http.StatusInternalServerError, err)
	}

	// TODO send email
}

func main() {
	lambda.Start(function.New(&function.Config{}, handler))
}
