package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/email"
	"github.com/g-harel/targetblank/internal/function"
	"github.com/g-harel/targetblank/internal/hash"
	"github.com/g-harel/targetblank/internal/page"
	"github.com/g-harel/targetblank/internal/rand"
	"github.com/g-harel/targetblank/internal/tables"
)

var messageSubject = "Your new homepage is ready!"

type messageContent struct {
	Addr  string
	Token string
}

var messageTemplate = `
	<html>
		<body>
			<h3>Follow the link to confirm you're the owner.</h3>
			<span>https://targetblank.org/{{.Addr}}/reset/{{.Token}}</span>
		</body>
	</html>`

var pages tables.IPage
var sender email.ISender

func init() {
	pages = tables.NewPage()
	sender = email.NewSender()
}

var defaultPage = "version 1\n==="

func handler(req *function.Request, res *function.Response) *function.Error {
	e := strings.TrimSpace(req.Body)

	match, err := regexp.MatchString(`^\S+@\S+\.\S+$`, e)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}
	if !match {
		return function.CustomErr(fmt.Errorf("invalid email address"))
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

	body, err := email.Template(messageTemplate, &messageContent{
		Addr:  item.Key,
		Token: token,
	})
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}

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
