package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/database"
	"github.com/g-harel/targetblank/internal/database/pages"
	"github.com/g-harel/targetblank/internal/function"
)

func handler(req *function.Request, res *function.Response) {
	item := &pages.Item{
		Password:           req.Body,
		TempPass:           false,
		TempPassHasBeenSet: true,
	}

	err := database.Validate(item.Password, "password")
	if err != nil {
		res.ClientErr(http.StatusBadRequest, err)
		return
	}
	item.Password, err = database.Hash(item.Password)
	if err != nil {
		res.ServerErr(http.StatusInternalServerError, err)
		return
	}

	err = pages.New(database.New()).Change(req.PathParameters["address"], item)
	if err != nil {
		res.ServerErr(http.StatusInternalServerError, err)
		return
	}
}

func main() {
	lambda.Start(function.New(&function.Config{
		PathParams: []string{"address"},
	}, handler))
}
