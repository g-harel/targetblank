package handler

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// Request replaces the api gateway request event.
type Request events.APIGatewayProxyRequest

// Param reads a path parameter from the request.
func (r *Request) Param(n string) (string, *Error) {
	if r.PathParameters == nil || r.PathParameters[n] == "" {
		return "", ClientErr("missing path param '%v'.", n)
	}
	return r.PathParameters[n], nil
}

// Response replaces the api gateway response event.
type Response events.APIGatewayProxyResponse

// ContentType adds the content-type header to the request.
func (r *Response) ContentType(t string) {
	if r.Headers == nil {
		r.Headers = map[string]string{}
	}
	r.Headers["Content-Type"] = t
}

// Handler is a custom type representing a lambda handler.
type Handler func(*Request, *Response) *Error

// New creates a lambda handler from a Handler and a Config.
func New(h Handler) func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		request := Request(req)

		// Default handler response has an empty JSON body and status code of 200.
		response := &Response{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			},
			Body: "{}",
		}
		response.ContentType("application/json")

		funcErr := h(&request, response)
		if funcErr != nil {
			response.StatusCode = funcErr.code
			response.ContentType("text/plain")
			response.Body = funcErr.Error()
		}

		return events.APIGatewayProxyResponse(*response), nil
	}
}
