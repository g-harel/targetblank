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

func handler(req *function.Request, res *function.Response) {
	item := &pages.Item{Email: req.Body}

	err := database.Validate(item.Email, "email")
	if err != nil {
		res.ClientErr(http.StatusBadRequest, err)
		return
	}
	email, err := database.Hash(item.Email)
	if err != nil {
		res.ServerErr(http.StatusInternalServerError, err)
		return
	}
	item.Email = email

	pass, err := database.Hash(database.RandString(16))
	if err != nil {
		res.ServerErr(http.StatusInternalServerError, err)
		return
	}
	item.Password = pass
	item.TempPass = true

	page, parseErr := page.NewFromSpec("version 1\n===")
	if parseErr != nil {
		res.ServerErr(http.StatusInternalServerError, parseErr)
		return
	}
	marshalledPageB, err := json.Marshal(page)
	if err != nil {
		res.ServerErr(http.StatusInternalServerError, err)
		return
	}
	item.Page = string(marshalledPageB)

	item.Published = false

	err = pages.New(database.New()).Create(item)
	if err != nil {
		res.ServerErr(http.StatusInternalServerError, err)
	}

	// TODO send email
}

func main() {
	lambda.Start(function.New(&function.Config{}, handler))
}
