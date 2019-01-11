package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/crypto"
	"github.com/g-harel/targetblank/internal/handlers"
	"github.com/g-harel/targetblank/services/storage"
)

var storagePageRead = storage.PageRead

func handler(req *handlers.Request, res *handlers.Response) *handlers.Error {
	addr, funcErr := req.Param("addr")
	if funcErr != nil {
		return funcErr
	}

	page, err := storagePageRead(addr)
	if err != nil {
		return handlers.InternalErr("read page: %v", err)
	}
	if page == nil {
		return handlers.ClientErr(handlers.ErrPageNotFound)
	}

	if !crypto.HashCheck(req.Body, page.Password) {
		return handlers.ClientErr(handlers.ErrPageNotFound)
	}

	token, err := handlers.CreateToken(false, addr)
	if err != nil {
		return handlers.InternalErr("create token: %v", err)
	}

	res.Body = token
	res.ContentType("text/plain")

	return nil
}

func main() {
	lambda.Start(handlers.New(handler))
}
