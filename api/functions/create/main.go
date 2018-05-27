package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/api/internal/email"
	"github.com/g-harel/targetblank/api/internal/function"
	"github.com/g-harel/targetblank/api/internal/hash"
	"github.com/g-harel/targetblank/api/internal/kind"
	"github.com/g-harel/targetblank/api/internal/page"
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
	sender = email.New()
}

var defaultPage = "version 1\n==="

func handler(req *function.Request, res *function.Response) *function.Error {
	e := strings.TrimSpace(req.Body)

	err := kind.Of(e).Is(kind.EMAIL)
	if err != nil {
		return function.CustomErr(err)
	}
	h, err := hash.New(e)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}
	item := &tables.PageItem{Email: h}

	pass, err := hash.New(rand.String(16))
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}
	item.Password = pass
	item.TempPass = true

	page, parseErr := page.NewFromSpec(defaultPage)
	if parseErr != nil {
		return function.Err(http.StatusInternalServerError, parseErr)
	}
	marshalledPageB, err := json.Marshal(page)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}
	item.Page = string(marshalledPageB)

	item.Published = false

	err = pages.Create(item)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}

	token, funcErr := function.MakeToken(true, item.Key)
	if funcErr != nil {
		return funcErr
	}

	body := strings.TrimSpace(messageContent)
	body = strings.Replace(body, "\n\t", "\n", -1)

	body = strings.Replace(body, "{{addr}}", item.Key, -1)
	body = strings.Replace(body, "{{token}}", token, -1)

	err = sender.Send(e, messageSubject, body)
	if err != nil {
		return function.Err(
			http.StatusInternalServerError,
			fmt.Errorf("Error sending email: %v", err),
		)
	}

	res.Body = item.Key
	res.ContentType("text/plain")

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
