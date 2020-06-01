package main

import (
	"crypto/rand"
	"math/big"
	"regexp"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/crypto"
	"github.com/g-harel/targetblank/internal/handler"
	"github.com/g-harel/targetblank/internal/parse"
	"github.com/g-harel/targetblank/services/mailer"
	"github.com/g-harel/targetblank/services/secrets"
	"github.com/g-harel/targetblank/services/storage"
)

var mailerSend = mailer.Send
var secretsKey = secrets.Key
var storagePageCreate = storage.PageCreate

var defaultDocument = `version 1
title = Welcome!
===
This is your new page, make sure you save the url so that you can come back to it.
You can also share the url with others if you want to show off your links!
Click the edit button in the top right to start making changes.
---
Helpful Links
    Document Format [https://github.com/g-harel/targetblank/#document-format]
    Keyboard Shortcuts [https://github.com/g-harel/targetblank/#keyboard-shortcuts]
`

// Generates a pseudorandom page id.
func genPageID() string {
	// Alphabet of unambiguous characters (without "Il0O").
	var alphabet = []rune("123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")

	b := make([]rune, 6)
	for i := range b {
		pos, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		if err != nil {
			panic(err)
		}
		b[i] = alphabet[pos.Int64()]
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

	page.Published = true

	// Loop until an available address is found.
	for {
		page.Addr = genPageID()
		err := storagePageCreate(page)
		if err == storage.ErrFailedCondition {
			// Try again.
			continue
		}
		if err != nil {
			return handler.InternalErr("create page: %v", err)
		}
		break
	}

	key, err := secretsKey()
	if err != nil {
		return handler.InternalErr("read secret key: %v", err)
	}

	token, err := handler.CreateToken(key, true, page.Addr)
	if err != nil {
		return handler.InternalErr("create token: %v", err)
	}

	err = mailerSend(
		email,
		"Your page is ready!",
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
