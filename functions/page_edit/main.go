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
	addr, err := req.Param("address")
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}

	_, funcErr := req.ValidateToken(addr)
	if err != nil {
		return funcErr
	}

	page, parseErr := page.NewFromSpec(req.Body)
	if parseErr != nil {
		return function.CustomErr(parseErr)
	}

	bytes, err := json.Marshal(page)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}
	item := &pages.Item{
		Page: string(bytes),
	}

	err = pages.New(client).Change(addr, item)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}

	res.Body = item.Page
	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
