package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/handler"
	"github.com/g-harel/targetblank/internal/parse"
)

// Validate parses the supplied document template.
func Validate(req *handler.Request, res *handler.Response) *handler.Error {
	doc, err := parse.Document(req.Body)
	if err != nil {
		return handler.ClientErr("parsing error: %v", err)
	}

	res.Body = doc

	return nil
}

func main() {
	lambda.Start(handler.New(Validate))
}
