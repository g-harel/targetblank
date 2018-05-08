package main

import (
	"encoding/json"
	"net/http"

	"github.com/g-harel/targetblank/internal/database"
	"github.com/g-harel/targetblank/internal/database/pages"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/targetblank/internal/function"
	"github.com/g-harel/targetblank/internal/page"
)

func handler(req *function.Request, res *function.Response) {
	page, parseErr := page.NewFromSpec(req.Body)
	if parseErr != nil {
		res.ClientErr(http.StatusBadRequest, parseErr)
		return
	}

	bytes, err := json.Marshal(page)
	if err != nil {
		res.ServerErr(http.StatusInternalServerError, err)
	}
	p := string(bytes)

	err = pages.New(database.New()).Edit(req.PathParameters["address"], pages.Item{
		Page: p,
	})
	if err != nil {
		res.ServerErr(http.StatusInternalServerError, err)
		return
	}

	res.Body = p
}

func main() {
	lambda.Start(function.New(&function.Config{
		RequireAuth: true,
		PathParams:  []string{"address"},
	}, handler))
}
