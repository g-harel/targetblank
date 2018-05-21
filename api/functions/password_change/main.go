package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/api/internal/database"
	"github.com/g-harel/targetblank/api/internal/database/pages"
	"github.com/g-harel/targetblank/api/internal/function"
	"github.com/g-harel/targetblank/api/internal/hash"
	"github.com/g-harel/targetblank/api/internal/kind"
)

var client = database.New()

func handler(req *function.Request, res *function.Response) *function.Error {
	addr, funcErr := req.Param("addr")
	if funcErr != nil {
		return funcErr
	}

	_, funcErr = req.ValidateToken(addr)
	if funcErr != nil {
		return funcErr
	}

	item := &pages.Item{
		Password:           req.Body,
		TempPass:           false,
		TempPassHasBeenSet: true,
	}

	err := kind.Of(item.Password).Is(kind.PASSWORD)
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
