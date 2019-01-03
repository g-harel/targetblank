package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/crypto"
	"github.com/g-harel/targetblank/internal/function"
	"github.com/g-harel/targetblank/internal/mailer"
	"github.com/g-harel/targetblank/internal/tables"
)

var pages tables.IPage
var mailerSend = mailer.Send

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

	email := strings.TrimSpace(req.Body)

	ok := crypto.HashCheck(email, item.Email)
	if !ok {
		return function.Err(http.StatusBadRequest, errors.New("email does not match hashed value"))
	}

	err = pages.Change(addr, &tables.PageItem{
		Password: "invalid password",
	})
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}

	token, funcErr := function.MakeToken(true, addr)
	if funcErr != nil {
		return funcErr
	}

	err = mailerSend(
		email,
		"Your homepage password has been reset!",
		`<html>
			<body>
				<h3>Follow the link to confirm you're the owner.</h3>
				<span>https://targetblank.org/{{.Addr}}/reset/{{.Token}}</span>
			</body>
		</html>`,
		&struct {
			Addr  string
			Token string
		}{
			Addr:  item.Key,
			Token: token,
		},
	)
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
