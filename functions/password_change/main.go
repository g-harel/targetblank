package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/database"
	"github.com/g-harel/targetblank/internal/database/pages"
	"github.com/g-harel/targetblank/internal/function"
	"github.com/g-harel/targetblank/internal/hash"
	"github.com/g-harel/targetblank/internal/kind"
)

var client = database.New()

func handler(req *function.Request, res *function.Response) *function.Error {
	addr, err := req.Param("address")
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}

	_, funcErr := req.ValidateToken(addr)
	if err != nil {
		return funcErr
	}

	item := &pages.Item{
		Password:           req.Body,
		TempPass:           false,
		TempPassHasBeenSet: true,
	}

	err = kind.Of(item.Password).Is(kind.PASSWORD)
	if err != nil {
		return function.CustomErr(err)
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
