package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/api/internal/function"
	"github.com/g-harel/targetblank/api/internal/hash"
	"github.com/g-harel/targetblank/api/internal/kind"
	"github.com/g-harel/targetblank/api/internal/page"
	"github.com/g-harel/targetblank/api/internal/rand"
	"github.com/g-harel/targetblank/api/internal/tables"
)

var pages = tables.NewPage()

var defaultPage = "version 1\n==="

func handler(req *function.Request, res *function.Response) *function.Error {
	item := &tables.PageItem{Email: req.Body}

	err := kind.Of(item.Email).Is(kind.EMAIL)
	if err != nil {
		return function.CustomErr(err)
	}
	email, err := hash.New(item.Email)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}
	item.Email = email

	pass, err := hash.New(rand.String(16))
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}
	item.Password = pass
	item.TempPass = true

	page, parseErr := page.NewFromSpec(defaultPage)
	if parseErr != nil {
		return function.Err(http.StatusInternalServerError, parseErr)
	}
	marshalledPageB, err := json.Marshal(page)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}
	item.Page = string(marshalledPageB)

	item.Published = false

	err = pages.Create(item)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}

	// TODO send email

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
