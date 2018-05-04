package function

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// Handler is a translated lambda handler.
type Handler func(*Request, *Response)

// Config contains options for the handler's middleware.
type Config struct{}

// Request replaces the api gateway request event.
type Request events.APIGatewayProxyRequest

// Response replaces the api gateway response event.
type Response events.APIGatewayProxyResponse

// ClientErr adds a client error to the response.
func (r *Response) ClientErr(status int, f string, args ...interface{}) {
	r.StatusCode = status
	r.Body = fmt.Sprintf(f, args...)
}

// ServerErr adds a server error to the response.
func (r *Response) ServerErr(status int, f string, args ...interface{}) {
	err := fmt.Sprintf(f, args...)
	fmt.Println(err) // TODO proper logging
	r.StatusCode = status
	r.Body = http.StatusText(status)
}

// New creates a lambda handler from a custom handler and a config.
func New(c *Config, h Handler) func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(req events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
		request := Request(req)
		response := &Response{
			StatusCode: http.StatusOK,
		}
		h(&request, response)
		return events.APIGatewayProxyResponse(*response), nil
	}
}
