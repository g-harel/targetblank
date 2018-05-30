package main

import (
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/api/internal/function"
	"github.com/g-harel/targetblank/api/internal/hash"
	"github.com/g-harel/targetblank/api/internal/kind"
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

	_, funcErr = req.ValidateToken(addr)
	if funcErr != nil {
		return funcErr
	}

	pass := strings.TrimSpace(req.Body)

	err := kind.Of(pass).Is(kind.PASSWORD)
	if err != nil {
		return function.CustomErr(err)
	}
	h, err := hash.New(pass)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}

	item := &tables.PageItem{
		Password: h,
	}

	err = pages.Change(addr, item)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
