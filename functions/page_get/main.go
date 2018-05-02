package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/page"
)

func echo() (*page.Page, error) {
	return page.New(), nil
}

func main() {
	lambda.Start(echo)
}
