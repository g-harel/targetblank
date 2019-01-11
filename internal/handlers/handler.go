package handlers

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// Request replaces the api gateway request event.
type Request events.APIGatewayProxyRequest

// Param reads a path parameter from the request.
func (r *Request) Param(n string) (string, *Error) {
	if r.PathParameters == nil || r.PathParameters[n] == "" {
		return "", InternalErr("read path param: %v", n)
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
		response := &Response{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Content-Type": "application/json",
				// CORS headers.
				"Access-Control-Allow-Origin":      "*",
				"Access-Control-Allow-Headers":     req.Headers["Access-Control-Request-Headers"],
				"Access-Control-Allow-Methods":     req.HTTPMethod,
				"Access-Control-Allow-Credentials": "true",
			},
			Body: "{}",
		}

		funcErr := h(&request, response)
		if funcErr != nil {
			response.StatusCode = funcErr.code
			response.ContentType("text/plain")
			response.Body = funcErr.Error()
		}

		return events.APIGatewayProxyResponse(*response), nil
	}
}
