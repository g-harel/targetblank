package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/function"
	"github.com/g-harel/targetblank/services/storage"
)

var storagePageRead = storage.PageRead

func handler(req *function.Request, res *function.Response) *function.Error {
	addr, funcErr := req.Param("addr")
	if funcErr != nil {
		return funcErr
	}

	page, err := storagePageRead(addr)
	if err != nil {
		return function.InternalErr("read page: %v", err)
	}
	if page == nil {
		return function.ClientErr("page not found")
	}

	if !page.Published {
		funcErr = req.Authenticate(addr)
		if funcErr != nil {
			// Page existence is kept hidden.
			return function.ClientErr("page not found")
		}
	}

	res.Body = page.Document

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
