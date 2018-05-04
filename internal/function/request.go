package function

import "github.com/aws/aws-lambda-go/events"

// Request replaces the api gateway request event.
type Request events.APIGatewayProxyRequest
