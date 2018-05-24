package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/api/internal/function"
	"github.com/g-harel/targetblank/api/internal/page"
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

	page, parseErr := page.NewFromSpec(req.Body)
	if parseErr != nil {
		return function.CustomErr(parseErr)
	}

	bytes, err := json.Marshal(page)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}
	item := &tables.PageItem{
		Page: string(bytes),
	}

	err = pages.Change(addr, item)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}

	res.Body = item.Page
	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
