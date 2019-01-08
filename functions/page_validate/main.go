package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/function"
	"github.com/g-harel/targetblank/internal/parser"
)

func handler(req *function.Request, res *function.Response) *function.Error {
	_, parseErr := parser.ParseDocument(req.Body)
	if parseErr != nil {
		return function.CustomErr(parseErr)
	}

	return nil
}

func main() {
	lambda.Start(function.New(handler))
}
