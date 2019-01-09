package main

import (
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/crypto"
	"github.com/g-harel/targetblank/internal/function"
	"github.com/g-harel/targetblank/services/mailer"
	"github.com/g-harel/targetblank/services/storage"
)

var mailerSend = mailer.Send
var storagePageRead = storage.PageRead

func handler(req *function.Request, res *function.Response) *function.Error {
	addr, funcErr := req.Param("addr")
	if funcErr != nil {
		return funcErr
	}

	page, err := storagePageRead(addr)
	if err != nil {
		return function.InternalErr("read page: %v", err)
	}
	if page == nil {
		return function.ClientErr("page not found")
	}

	email := strings.TrimSpace(req.Body)

	ok := crypto.HashCheck(email, page.Email)
	if !ok {
		return function.ClientErr("page not found")
	}

	token, err := function.CreateToken(true, addr)
	if err != nil {
		return function.InternalErr("create token: %v", err)
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
			Addr:  page.Addr,
			Token: token,
		},
	)
	if err != nil {
		return function.InternalErr("send email: %v", err)
	}

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
