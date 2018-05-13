package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/database"
	"github.com/g-harel/targetblank/internal/database/pages"
	"github.com/g-harel/targetblank/internal/function"
	"github.com/g-harel/targetblank/internal/hash"
	"github.com/g-harel/targetblank/internal/rand"
)

var client = database.New()

func handler(req *function.Request, res *function.Response) *function.Error {
	addr, err := req.Param("address")
	if err != nil {
		return function.ServerErr(http.StatusInternalServerError, err)
	}

	pass := rand.String(16)
	item := &pages.Item{
		TempPass:           true,
		TempPassHasBeenSet: true,
	}

	item.Password, err = hash.New(pass)
	if err != nil {
		return function.ServerErr(http.StatusInternalServerError, err)
	}

	err = pages.New(client).Change(addr, item)
	if err != nil {
		return function.ServerErr(http.StatusInternalServerError, err)
	}

	// TODO send email

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
