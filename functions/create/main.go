package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/crypto"
	"github.com/g-harel/targetblank/internal/function"
	"github.com/g-harel/targetblank/internal/mailer"
	"github.com/g-harel/targetblank/internal/page"
	"github.com/g-harel/targetblank/storage"
)

var pages storage.IPage
var mailerSend = mailer.Send

func init() {
	pages = storage.NewPage()
}

var defaultPage = "version 1\n==="

func handler(req *function.Request, res *function.Response) *function.Error {
	email := strings.TrimSpace(req.Body)

	match, err := regexp.MatchString(`^\S+@\S+\.\S+$`, email)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}
	if !match {
		return function.CustomErr(fmt.Errorf("invalid email address"))
	}
	emailHash, err := crypto.Hash(email)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}
	item := &storage.PageItem{Email: emailHash}

	pass := make([]byte, 16)
	_, err = rand.Read(pass)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}

	passHash, err := crypto.Hash(string(pass))
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}
	item.Password = passHash

	page, parseErr := page.NewFromSpec(defaultPage)
	if parseErr != nil {
		return function.Err(http.StatusInternalServerError, parseErr)
	}
	marshalledPage, err := json.Marshal(page)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}
	item.Page = string(marshalledPage)

	item.Published = false

	err = pages.Create(item)
	if err != nil {
		return function.Err(http.StatusInternalServerError, err)
	}

	token, funcErr := function.MakeToken(true, item.Key)
	if funcErr != nil {
		return funcErr
	}

	err = mailerSend(
		email,
		"Your new homepage is ready!",
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

	res.Body = item.Key
	res.ContentType("text/plain")

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
