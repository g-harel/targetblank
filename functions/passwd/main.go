package main

import (
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/crypto"
	"github.com/g-harel/targetblank/internal/handler"
	"github.com/g-harel/targetblank/services/secrets"
	"github.com/g-harel/targetblank/services/storage"
)

var secretsKey = secrets.Key
var storagePageUpdatePassword = storage.PageUpdatePassword

// Passwd updates the page password.
func Passwd(req *handler.Request, res *handler.Response) *handler.Error {
	addr, funcErr := req.Param("addr")
	if funcErr != nil {
		return funcErr
	}

	key, err := secretsKey()
	if err != nil {
		return handler.InternalErr("read secret key: %v", err)
	}

	_, funcErr = req.Authenticate(key, addr)
	if funcErr != nil {
		return funcErr
	}

	pass := strings.TrimSpace(req.Body)

	if len(pass) < 8 {
		return handler.ClientErr(handler.ErrInvalidPassword)
	}
	h, err := crypto.Hash(pass)
	if err != nil {
		return handler.InternalErr("hash password: %v", err)
	}

	err = storagePageUpdatePassword(addr, h)
	if err != nil {
		return handler.InternalErr("update page password: %v", err)
	}

	return nil
}

func main() {
	lambda.Start(handler.New(Passwd))
}
