package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/function"
	"github.com/g-harel/targetblank/internal/parser"
	"github.com/g-harel/targetblank/services/storage"
)

var storagePageUpdateDocument = storage.PageUpdateDocument

func handler(req *function.Request, res *function.Response) *function.Error {
	addr, funcErr := req.Param("addr")
	if funcErr != nil {
		return funcErr
	}

	_, funcErr = req.ValidateToken(addr)
	if funcErr != nil {
		return funcErr
	}

	page, parseErr := parser.ParseDocument(req.Body)
	if parseErr != nil {
		return function.CustomErr(parseErr)
	}

	bytes, err := json.Marshal(page)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}

	err = storagePageUpdateDocument(addr, string(bytes))
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}

	res.Body = string(bytes)

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
