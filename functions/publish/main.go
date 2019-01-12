package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/handler"
	"github.com/g-harel/targetblank/services/storage"
)

var storagePageUpdatePublished = storage.PageUpdatePublished

// Publish makes the page readable without credentials.
func Publish(req *handler.Request, res *handler.Response) *handler.Error {
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
		return handler.InternalErr("update page published: %v", err)
	}

	return nil
}

func main() {
	lambda.Start(handler.New(Publish))
}
