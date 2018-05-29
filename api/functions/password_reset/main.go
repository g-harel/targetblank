package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/g-harel/targetblank/api/internal/email"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/api/internal/function"
	"github.com/g-harel/targetblank/api/internal/hash"
	"github.com/g-harel/targetblank/api/internal/rand"
	"github.com/g-harel/targetblank/api/internal/tables"
)

var messageSubject = "Your new homepage is ready!"
var messageContent = `
	addr: {{addr}}
	token: {{token}}
`

var pages tables.IPage
var sender email.ISender

func init() {
	pages = tables.NewPage()
	sender = email.NewSender()
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

	email := strings.TrimSpace(req.Body)

	ok := hash.Check(email, item.Email)
	if !ok {
		return function.Err(http.StatusBadRequest, errors.New("email does not match hashed value"))
	}

	h, err := hash.New(rand.String(16))
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}

	err = pages.Change(addr, &tables.PageItem{
		Password: h,
	})
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}

	token, funcErr := function.MakeToken(true, addr)
	if funcErr != nil {
		return funcErr
	}

	body := strings.TrimSpace(messageContent)
	body = strings.Replace(body, "\n\t", "\n", -1)

	body = strings.Replace(body, "{{addr}}", item.Key, -1)
	body = strings.Replace(body, "{{token}}", token, -1)

	err = sender.Send(email, messageSubject, body)
	if err != nil {
		return function.Err(
			http.StatusInternalServerError,
			fmt.Errorf("Error sending email: %v", err),
		)
	}

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
