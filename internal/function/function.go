package function

import (
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// Config contains options for the handler's middleware.
type Config struct {
	PathParams []string
}

// Handler is a custom type representing a lambda handler.
type Handler func(*Request, *Response)

// New creates a lambda handler from a Handler and a Config.
func New(c *Config, h Handler) func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// TODO read jwt
	return func(req events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
		request := Request(req)
		response := &Response{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: "{}",
		}

		if c.PathParams != nil {
			for _, p := range c.PathParams {
				if request.PathParameters[p] == "" {
					response.ClientErr(http.StatusBadRequest, errors.New("missing path parameter: "+p))
				}
			}
		}

		h(&request, response)
		return events.APIGatewayProxyResponse(*response), nil
	}
}
