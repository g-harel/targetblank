package main

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	resp, err := HandleRequest(events.APIGatewayProxyRequest{})
	if err != nil {
		t.Fatalf("Error in request handler: %v", err)
	}

	fmt.Println(resp.Body)

	HandleRequest(events.APIGatewayProxyRequest{})
	HandleRequest(events.APIGatewayProxyRequest{})
	HandleRequest(events.APIGatewayProxyRequest{})
	resp, _ = HandleRequest(events.APIGatewayProxyRequest{})
	fmt.Println(resp.Body)
}
