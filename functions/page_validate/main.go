package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/handlers"
	"github.com/g-harel/targetblank/internal/parse"
)

func handler(req *handlers.Request, res *handlers.Response) *handlers.Error {
	_, err := parse.Document(req.Body)
	if err != nil {
		return handlers.ClientErr("parsing error: %v", err)
	}

	return nil
}

func main() {
	lambda.Start(handlers.New(handler))
}
