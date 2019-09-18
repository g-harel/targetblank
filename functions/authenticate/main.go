package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/crypto"
	"github.com/g-harel/targetblank/internal/handler"
	"github.com/g-harel/targetblank/services/secrets"
	"github.com/g-harel/targetblank/services/storage"
)

var secretsKey = secrets.Key
var storagePageRead = storage.PageRead

// Authenticate responds with a token when given valid credentials.
func Authenticate(req *handler.Request, res *handler.Response) *handler.Error {
	addr, funcErr := req.Param("addr")
	if funcErr != nil {
		return funcErr
	}

	page, err := storagePageRead(addr)
	if err != nil {
		return handler.InternalErr("read page: %v", err)
	}
	if page == nil {
		return handler.ClientErr(handler.ErrPageNotFound)
	}

	if !crypto.HashCheck(req.Body, page.Password) {
		// Page existence is kept hidden.
		return handler.ClientErr(handler.ErrPageNotFound)
	}

	key, err := secretsKey()
	if err != nil {
		return handler.InternalErr("read secret key: %v", err)
	}

	token, err := handler.CreateToken(key, false, addr)
	if err != nil {
		return handler.InternalErr("create token: %v", err)
	}

	res.Body = token
	res.ContentType("text/plain")

	return nil
}

func main() {
	lambda.Start(handler.New(Authenticate))
}
