package main

import (
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/api/internal/function"
	"github.com/g-harel/targetblank/api/internal/tables"
)

var pages tables.IPage

func init() {
	pages = tables.NewPage()
}

func handler(req *function.Request, res *function.Response) *function.Error {
	addr, funcErr := req.Param("addr")
	if funcErr != nil {
		return funcErr
	}

	item, err := pages.Fetch(addr)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}
	if item == nil {
		return function.Err(http.StatusBadRequest, errors.New("page not found for given key"))
	}

	if !item.Published {
		_, funcErr = req.ValidateToken(addr)
		if funcErr != nil {
			return funcErr
		}
	}

	res.Body = item.Page

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
