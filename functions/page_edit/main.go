package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/handlers"
	"github.com/g-harel/targetblank/internal/parse"
	"github.com/g-harel/targetblank/services/storage"
)

var storagePageUpdateDocument = storage.PageUpdateDocument

func handler(req *handlers.Request, res *handlers.Response) *handlers.Error {
	addr, funcErr := req.Param("addr")
	if funcErr != nil {
		return funcErr
	}

	funcErr = req.Authenticate(addr)
	if funcErr != nil {
		return funcErr
	}

	page, err := parse.Document(req.Body)
	if err != nil {
		return handlers.ClientErr("parsing error: %v", err)
	}

	err = storagePageUpdateDocument(addr, page)
	if err != nil {
		return handlers.InternalErr("update page document: %v", err)
	}

	res.Body = page

	return nil
}

func main() {
	lambda.Start(handlers.New(handler))
}
