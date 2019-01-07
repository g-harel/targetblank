package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/crypto"
	"github.com/g-harel/targetblank/internal/function"
	"github.com/g-harel/targetblank/internal/mailer"
	"github.com/g-harel/targetblank/internal/page"
	"github.com/g-harel/targetblank/storage"
)

var mailerSend = mailer.Send
var storagePageCreate = storage.PageCreate

var defaultPage = "version 1\n==="

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generates a pseudorandom page id.
func genPageID() string {
	// List of unambiguous characters (minus "Il0O").
	var alphabet = []rune("123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")

	b := make([]rune, 6)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}

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
	item := &storage.Page{Email: emailHash}

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
	item.Data = string(marshalledPage)

	item.Published = false

	// Loop until an available address is found.
	for {
		item.Addr = genPageID()
		conflict, err := storagePageCreate(item)
		if err != nil {
			return function.Err(http.StatusInternalServerError, err)
		}
		if !conflict {
			break
		}
	}

	token, funcErr := function.MakeToken(true, item.Addr)
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
			Addr:  item.Addr,
			Token: token,
		},
	)
	if err != nil {
		return function.Err(
			http.StatusInternalServerError,
			fmt.Errorf("Error sending email: %v", err),
		)
	}

	res.Body = item.Addr
	res.ContentType("text/plain")

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
