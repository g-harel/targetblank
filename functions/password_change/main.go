package main

import (
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/crypto"
	"github.com/g-harel/targetblank/internal/function"
	"github.com/g-harel/targetblank/services/storage"
)

var storagePageUpdatePassword = storage.PageUpdatePassword

func handler(req *function.Request, res *function.Response) *function.Error {
	addr, funcErr := req.Param("addr")
	if funcErr != nil {
		return funcErr
	}

	funcErr = req.Authenticate(addr)
	if funcErr != nil {
		return funcErr
	}

	pass := strings.TrimSpace(req.Body)

	if len(pass) < 8 {
		return function.ClientErr("password is too short")
	}
	h, err := crypto.Hash(pass)
	if err != nil {
		return function.InternalErr("hash password: %v", err)
	}

	err = storagePageUpdatePassword(addr, h)
	if err != nil {
		return function.InternalErr("update page password: %v", err)
	}

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
