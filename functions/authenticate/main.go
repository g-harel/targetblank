package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/crypto"
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

	if !crypto.HashCheck(req.Body, page.Password) {
		return function.ClientErr("page not found")
	}

	token, err := function.CreateToken(false, addr)
	if err != nil {
		return function.InternalErr("create token: %v", err)
	}

	res.Body = token
	res.ContentType("text/plain")

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
