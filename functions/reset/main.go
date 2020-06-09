package main

import (
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/crypto"
	"github.com/g-harel/targetblank/internal/handler"
	"github.com/g-harel/targetblank/services/mailer"
	"github.com/g-harel/targetblank/services/secrets"
	"github.com/g-harel/targetblank/services/storage"
)

var secretsKey = secrets.Key
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
		// Page existence is kept hidden.
		return nil
	}

	email := strings.TrimSpace(req.Body)

	ok := crypto.HashCheck(email, page.Email)
	if !ok {
		// Configured email cannot be verified.
		return nil
	}

	key, err := secretsKey()
	if err != nil {
		return handler.InternalErr("read secret key: %v", err)
	}

	token, err := handler.CreateRestrictedToken(key, addr)
	if err != nil {
		return handler.InternalErr("create restricted token: %v", err)
	}

	err = mailerSend(
		email,
		"Reset your page password",
		`<html>
			<body>
				<p>
					<a href="https://targetblank.org">
						<img src="https://svgsaur.us/?t=targetblank&o=b&s=26&c=332832&w=152&y=55" />
						<br>
						<img src="https://svgsaur.us/?t=organize_your_links&s=16&c=766873&w=152&y=12" />
					</a>
				</p>
				<h3>
					<a href="https://targetblank.org/{{.Addr}}/reset/{{.Token}}">
						Update your password.
					</a>
				</h3>
				<p>
					<sup>Link will expire in 10 minutes.</sup>
				</p>
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
