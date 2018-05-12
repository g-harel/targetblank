package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/database"
	"github.com/g-harel/targetblank/internal/database/pages"
	"github.com/g-harel/targetblank/internal/function"
)

func handler(req *function.Request, res *function.Response) {
	pass := database.RandString(16)
	item := &pages.Item{
		TempPass:           true,
		TempPassHasBeenSet: true,
	}

	var err error
	item.Password, err = database.Hash(pass)
	if err != nil {
		res.ServerErr(http.StatusInternalServerError, err)
		return
	}

	err = pages.New(database.New()).Change(req.PathParameters["address"], item)
	if err != nil {
		res.ServerErr(http.StatusInternalServerError, err)
	}

	// TODO send email
}

func main() {
	lambda.Start(function.New(&function.Config{
		PathParams: []string{"address"},
	}, handler))
}
