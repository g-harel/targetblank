package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/function"
	"github.com/g-harel/targetblank/internal/parse"
)

func handler(req *function.Request, res *function.Response) *function.Error {
	_, err := parse.Document(req.Body)
	if err != nil {
		return function.ClientErr("parsing error: %v", err)
	}

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
