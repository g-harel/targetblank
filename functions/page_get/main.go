package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/g-harel/targetblank/internal/database"
	"github.com/g-harel/targetblank/internal/function"
)

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

func handler(req *function.Request, res *function.Response) {
	addr := req.Body
	input := &dynamodb.GetItemInput{
		TableName: aws.String("targetblank-pages"),
		Key: map[string]*dynamodb.AttributeValue{
			"addr": {
				S: aws.String(addr),
			},
		},
	}

	result, err := db.GetItem(input)
	if err != nil {
		res.ServerErr(http.StatusInternalServerError, err.Error())
		return
	}
	if result == nil {
		res.ClientErr(http.StatusNotFound, "page not found for key \"%v\"", addr)
	}

	item := &database.Item{}
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		res.ServerErr(http.StatusInternalServerError, "couldn't unmarshall response: %v", err)
		return
	}

	res.Body = item.Page
}

func main() {
	lambda.Start(function.New(&function.Config{}, handler))
}
