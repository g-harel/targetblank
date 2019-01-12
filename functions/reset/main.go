package main

import (
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/crypto"
	"github.com/g-harel/targetblank/internal/handler"
	"github.com/g-harel/targetblank/services/mailer"
	"github.com/g-harel/targetblank/services/storage"
)

var mailerSend = mailer.Send
var storagePageRead = storage.PageRead

// Reset sends a temporary password reset link to the page owner.
// Email is only sent if the supplied address is correct.
func Reset(req *handler.Request, res *handler.Response) *handler.Error {
	addr, funcErr := req.Param("addr")
	if funcErr != nil {
		return funcErr
	}

	page, err := storagePageRead(addr)
	if err != nil {
		return handler.InternalErr("read page: %v", err)
	}
	if page == nil {
		return handler.ClientErr(handler.ErrPageNotFound)
	}

	email := strings.TrimSpace(req.Body)

	ok := crypto.HashCheck(email, page.Email)
	if !ok {
		return handler.ClientErr(handler.ErrPageNotFound)
	}

	token, err := handler.CreateToken(true, addr)
	if err != nil {
		return handler.InternalErr("create token: %v", err)
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
		return handler.InternalErr("send email: %v", err)
	}

	return nil
}

func main() {
	lambda.Start(handler.New(Reset))
}
