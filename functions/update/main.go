package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/handler"
	"github.com/g-harel/targetblank/internal/parse"
	"github.com/g-harel/targetblank/services/storage"
)

var storagePageUpdateDocument = storage.PageUpdateDocument

// Update overrides the page document.
func Update(req *handler.Request, res *handler.Response) *handler.Error {
	addr, funcErr := req.Param("addr")
	if funcErr != nil {
		return funcErr
	}

	funcErr = req.Authenticate(addr)
	if funcErr != nil {
		return funcErr
	}

	doc, err := parse.Document(req.Body)
	if err != nil {
		return handler.ClientErr("parsing error: %v", err)
	}

	err = storagePageUpdateDocument(addr, doc)
	if err != nil {
		return handler.InternalErr("update page document: %v", err)
	}

	res.Body = doc
	res.ContentType("application/json")

	return nil
}

func main() {
	lambda.Start(handler.New(Update))
}
