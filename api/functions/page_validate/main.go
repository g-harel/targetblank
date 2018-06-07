package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/api/internal/function"
	"github.com/g-harel/targetblank/api/internal/page"
)

func handler(req *function.Request, res *function.Response) *function.Error {
	page, parseErr := page.NewFromSpec(req.Body)
	if parseErr != nil {
		return function.CustomErr(parseErr)
	}

	_, err := json.Marshal(page)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}