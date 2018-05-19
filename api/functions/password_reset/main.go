package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/api/internal/database"
	"github.com/g-harel/targetblank/api/internal/database/pages"
	"github.com/g-harel/targetblank/api/internal/function"
	"github.com/g-harel/targetblank/api/internal/hash"
	"github.com/g-harel/targetblank/api/internal/rand"
)

var client = database.New()

func handler(req *function.Request, res *function.Response) *function.Error {
	addr, funcErr := req.Param("address")
	if funcErr != nil {
		return funcErr
	}

	_, funcErr = req.ValidateToken(addr)
	if funcErr != nil {
		return funcErr
	}

	pass := rand.String(16)
	item := &pages.Item{
		TempPass:           true,
		TempPassHasBeenSet: true,
	}

	var err error
	item.Password, err = hash.New(pass)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}

	err = pages.New(client).Change(addr, item)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}

	// TODO send email

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
