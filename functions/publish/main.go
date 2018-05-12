package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/database"
	"github.com/g-harel/targetblank/internal/database/pages"
	"github.com/g-harel/targetblank/internal/function"
)

func handler(req *function.Request, res *function.Response) {
	err := pages.New(database.New()).Change(
		req.PathParameters["address"],
		pages.Item{
			Published:           true,
			PublishedHasBeenSet: true,
		},
	)
	if err != nil {
		res.ServerErr(http.StatusInternalServerError, err)
		return
	}
}

func main() {
	lambda.Start(function.New(&function.Config{
		PathParams: []string{"address"},
	}, handler))
}
