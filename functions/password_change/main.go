package main

import (
	"net/http"

	"github.com/g-harel/targetblank/internal/check"

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

	item := &pages.Item{
		Password:           req.Body,
		TempPass:           false,
		TempPassHasBeenSet: true,
	}

	err = check.That(item.Password).Is(check.PASSWORD)
	if err != nil {
		return function.CustomErr(http.StatusBadRequest, err)
	}
	item.Password, err = hash.New(item.Password)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}

	err = pages.New(client).Change(addr, item)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
