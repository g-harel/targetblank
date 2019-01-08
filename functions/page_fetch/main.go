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

	page, err := storagePageRead(addr)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}
	if page == nil {
		return function.Err(http.StatusBadRequest, errors.New("page not found for given address"))
	}

	if !page.Published {
		_, funcErr = req.ValidateToken(addr)
		if funcErr != nil {
			return funcErr
		}
	}

	res.Body = page.Document

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
