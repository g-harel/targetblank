package main

import (
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/crypto"
	"github.com/g-harel/targetblank/internal/handler"
	"github.com/g-harel/targetblank/internal/parse"
	"github.com/g-harel/targetblank/services/mailer"
	"github.com/g-harel/targetblank/services/storage"
)

var mailerSend = mailer.Send
var storagePageCreate = storage.PageCreate

var defaultDocument = "version 1\n==="

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generates a pseudorandom page id.
func genPageID() string {
	// Alphabet of unambiguous characters (without "Il0O").
	var alphabet = []rune("123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")

	b := make([]rune, 6)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}

// Create creates a new page for the given email address.
// An email is then sent to the owner with a temporary link.
func Create(req *handler.Request, res *handler.Response) *handler.Error {
	email := strings.TrimSpace(req.Body)

	match, err := regexp.MatchString(`^\S+@\S+\.\S+$`, email)
	if err != nil {
		return handler.InternalErr("match email pattern: %v", err)
	}
	if !match {
		return handler.ClientErr(handler.ErrInvalidEmail)
	}
	emailHash, err := crypto.Hash(email)
	if err != nil {
		return handler.InternalErr("hash email: %v", err)
	}
	page := &storage.Page{Email: emailHash}

	pass := make([]byte, 16)
	_, err = rand.Read(pass)
	if err != nil {
		return handler.InternalErr("generate random password: %v", err)
	}

	passHash, err := crypto.Hash(string(pass))
	if err != nil {
		return handler.InternalErr("hash password: %v", err)
	}
	page.Password = passHash

	doc, err := parse.Document(defaultDocument)
	if err != nil {
		return handler.InternalErr("parse default document: %v", err)
	}
	page.Document = doc

	page.Published = false

	// Loop until an available address is found.
	for {
		page.Addr = genPageID()
		conflict, err := storagePageCreate(page)
		if err != nil {
			return handler.InternalErr("create page: %v", err)
		}
		if !conflict {
			break
		}
	}

	token, err := handler.CreateToken(true, page.Addr)
	if err != nil {
		return handler.InternalErr("create token: %v", err)
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
			Addr:  page.Addr,
			Token: token,
		},
	)
	if err != nil {
		return handler.InternalErr("send email: %v", err)
	}

	res.Body = page.Addr
	res.ContentType("text/plain")

	return nil
}

func main() {
	lambda.Start(handler.New(Create))
}
