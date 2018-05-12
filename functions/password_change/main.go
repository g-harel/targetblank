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
		return function.ServerErr(http.StatusInternalServerError, err)
	}

	item := &pages.Item{
		Password:           req.Body,
		TempPass:           false,
		TempPassHasBeenSet: true,
	}

	err = database.Validate(item.Password, "password")
	if err != nil {
		return function.ClientErr(http.StatusBadRequest, err)
	}
	item.Password, err = database.Hash(item.Password)
	if err != nil {
		return function.ServerErr(http.StatusInternalServerError, err)
	}

	err = pages.New(client).Change(addr, item)
	if err != nil {
		return function.ServerErr(http.StatusInternalServerError, err)
	}

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
