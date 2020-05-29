package main

import (
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/handler"
	"github.com/g-harel/targetblank/services/secrets"
	"github.com/g-harel/targetblank/services/storage"
)

var secretsKey = secrets.Key
var storagePageRead = storage.PageRead

// Read responds with a parsed version of the page document.
// Authentication token is required if the page is not published.
func Read(req *handler.Request, res *handler.Response) *handler.Error {
	addr, funcErr := req.Param("addr")
	if funcErr != nil {
		return funcErr
	}

	page, err := storagePageRead(addr)
	if err != nil {
		return handler.InternalErr("read page: %v", err)
	}
	if page == nil {
		return handler.ObfuscatedAuthErr()
	}

	if !page.Published {
		key, err := secretsKey()
		if err != nil {
			return handler.InternalErr("read secret key: %v", err)
		}

		authTimestamp, funcErr := req.Authenticate(key, addr)
		if funcErr != nil {
			return handler.ObfuscatedAuthErr()
		}

		if page.PasswordLastUpdate != "" {
			passwordLastUpdate, err := time.Parse(storage.ISO8601, page.PasswordLastUpdate)
			if err != nil {
				return handler.InternalErr("parse 'PasswordLastUpdate' timestamp: %v", err)
			}
			if passwordLastUpdate.After(*authTimestamp) {
				return handler.ObfuscatedAuthErr()
			}
		}
	}

	res.Body = page.Document
	res.ContentType("application/json")

	return nil
}

func main() {
	lambda.Start(handler.New(Read))
}
