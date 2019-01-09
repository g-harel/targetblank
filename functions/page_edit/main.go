package main

import (
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

	funcErr = req.Authenticate(addr)
	if funcErr != nil {
		return funcErr
	}

	page, parseErr := parser.ParseDocument(req.Body)
	if parseErr != nil {
		return function.ClientErr("parsing error: %v", parseErr)
	}

	err := storagePageUpdateDocument(addr, page)
	if err != nil {
		return function.InternalErr("update page document: %v", err)
	}

	res.Body = page

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
