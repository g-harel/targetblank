package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/function"
	"github.com/g-harel/targetblank/services/storage"
)

var storagePageUpdatePublished = storage.PageUpdatePublished

func handler(req *function.Request, res *function.Response) *function.Error {
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
		return function.InternalErr("update page published: %v", err)
	}

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
