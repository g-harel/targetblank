package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/handlers"
	"github.com/g-harel/targetblank/services/storage"
)

var storagePageUpdatePublished = storage.PageUpdatePublished

func handler(req *handlers.Request, res *handlers.Response) *handlers.Error {
	addr, funcErr := req.Param("addr")
	if funcErr != nil {
		return funcErr
	}

	funcErr = req.Authenticate(addr)
	if funcErr != nil {
		return funcErr
	}

	err := storagePageUpdatePublished(addr, true)
	if err != nil {
		return handlers.InternalErr("update page published: %v", err)
	}

	return nil
}

func main() {
	lambda.Start(handlers.New(handler))
}
