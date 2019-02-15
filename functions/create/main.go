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
	"github.com/g-harel/targetblank/services/storage"
)

var mailerSend = mailer.Send
var storagePageCreate = storage.PageCreate

var defaultDocument = `# Everything after a pound character (#), trailing whitespace and empty lines are ignored.

# Documents must start with their version (currently only 1).
version 1

# Document metadata key-value pairs can be added at the top of the document.
key=value

# The "title" key can be used to name the document.
title=Hello World

# The first group starts after the header line.
===

# Group metadata key-value pairs can be added at the start of each group.
# These values are currently ignored, but may be used in the future.
key=value

# Groups hold entries containing a label and a link.
labelled link [example.com]

# Both the label and the link are optional.
label without link
[google.com]
amazon.com

# New groups are started using the group delimiter.
---

# Group entries can be nested using indentation.
entry 1
    entry 2
        entry 3
    entry 4
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
