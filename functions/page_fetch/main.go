package main

import (
	"errors"
	"net/http"

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

	item, err := storagePageRead(addr)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}
	if item == nil {
		return function.Err(http.StatusBadRequest, errors.New("page not found for given address"))
	}

	if !item.Published {
		_, funcErr = req.ValidateToken(addr)
		if funcErr != nil {
			return funcErr
		}
	}

	res.Body = item.Data

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
