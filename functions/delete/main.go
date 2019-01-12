package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/handler"
	"github.com/g-harel/targetblank/services/storage"
)

var storagePageDelete = storage.PageDelete

// Delete permanently deletes a page.
func Delete(req *handler.Request, res *handler.Response) *handler.Error {
	addr, funcErr := req.Param("addr")
	if funcErr != nil {
		return funcErr
	}

	funcErr = req.Authenticate(addr)
	if funcErr != nil {
		return funcErr
	}

	err := storagePageDelete(addr)
	if err != nil {
		return handler.InternalErr("delete page: %v", err)
	}

	return nil
}

func main() {
	lambda.Start(handler.New(Delete))
}
