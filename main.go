package main

import (
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var count = 0

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	count++
	return events.APIGatewayProxyResponse{
		Body:       strconv.Itoa(count),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
