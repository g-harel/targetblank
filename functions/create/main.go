package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/database"
	"github.com/g-harel/targetblank/internal/database/pages"
	"github.com/g-harel/targetblank/internal/function"
	"github.com/g-harel/targetblank/internal/page"
)

var client = database.New()

func handler(req *function.Request, res *function.Response) *function.Error {
	item := &pages.Item{Email: req.Body}

	err := database.Validate(item.Email, "email")
	if err != nil {
		return function.ClientErr(http.StatusBadRequest, err)
	}
	email, err := database.Hash(item.Email)
	if err != nil {
		return function.ServerErr(http.StatusInternalServerError, err)
	}
	item.Email = email

	pass, err := database.Hash(database.RandString(16))
	if err != nil {
		return function.ServerErr(http.StatusInternalServerError, err)
	}
	item.Password = pass
	item.TempPass = true

	page, parseErr := page.NewFromSpec("version 1\n===")
	if parseErr != nil {
		return function.ServerErr(http.StatusInternalServerError, parseErr)
	}
	marshalledPageB, err := json.Marshal(page)
	if err != nil {
		return function.ServerErr(http.StatusInternalServerError, err)
	}
	item.Page = string(marshalledPageB)

	item.Published = false

	err = pages.New(client).Create(item)
	if err != nil {
		return function.ServerErr(http.StatusInternalServerError, err)
	}

	// TODO send email

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
