package function

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// Request replaces the api gateway request event.
type Request events.APIGatewayProxyRequest

// Param reads a path parameter from the request.
func (r *Request) Param(n string) (string, error) {
	if r.PathParameters == nil || r.PathParameters[n] == "" {
		return "", fmt.Errorf("Could not access path parameter \"%v\"", n)
	}

	return r.PathParameters[n], nil
}
