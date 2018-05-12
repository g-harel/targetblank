package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/database"
	"github.com/g-harel/targetblank/internal/database/pages"
	"github.com/g-harel/targetblank/internal/function"
)

func handler(req *function.Request, res *function.Response) {
	err := pages.New(database.New()).Change(
		req.PathParameters["address"],
		pages.Item{
			Password:           req.Body,
			TempPass:           false,
			TempPassHasBeenSet: true,
		},
	)
	switch err.(type) {
	case nil:
	case database.ValidationError:
		res.ClientErr(http.StatusBadRequest, err)
	default:
		res.ServerErr(http.StatusInternalServerError, err)
	}
}

func main() {
	lambda.Start(function.New(&function.Config{
		PathParams: []string{"address"},
	}, handler))
}
