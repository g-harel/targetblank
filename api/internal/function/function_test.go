package function

import (
	"errors"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestRequestParam(t *testing.T) {
	t.Run("should return the request's path parameter", func(t *testing.T) {
		name := "param"
		value := "value"
		req := Request{
			PathParameters: map[string]string{
				name: value,
			},
		}

		v, err := req.Param(name)
		if err != nil {
			t.Fatal("Fetching existing path param produced error")
		}
		if v != value {
			t.Fatal("Fetched parameter has unexpected value")
		}
	})

	t.Run("should return an error if the parameter is missing", func(t *testing.T) {
		name := "param name"
		req := Request{}

		_, err := req.Param(name)
		if err == nil {
			t.Fatal("Empty PathParameters map should produce an error")
		}

		req.PathParameters = map[string]string{}
		_, err = req.Param(name)
		if err == nil {
			t.Fatal("Empty PathParameters value should produce an error")
		}
	})
}

func TestNew(t *testing.T) {
	t.Run("should produce event handlers compatible with lambda/api-gateway", func(t *testing.T) {
		var h Handler = func(req *Request, res *Response) *Error { return nil }
		var _ func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) = New(h)
	})

	t.Run("should run the handler", func(t *testing.T) {
		check := false
		var h Handler = func(req *Request, res *Response) *Error {
			check = true
			return nil
		}
		New(h)(events.APIGatewayProxyRequest{})
		if !check {
			t.Fatalf("Handler not being run by returned func")
		}
	})

	t.Run("should set default status to 200", func(t *testing.T) {
		var h Handler = func(req *Request, res *Response) *Error {
			if res.StatusCode != 200 {
				t.Fatal("Expected default statuscode to be 200")
			}
			return nil
		}
		New(h)(events.APIGatewayProxyRequest{})
	})

	t.Run("should set default response to empty json", func(t *testing.T) {
		var h Handler = func(req *Request, res *Response) *Error {
			if res.Headers["Content-Type"] != "application/json" {
				t.Fatal("Expected default content type to be json")
			}
			if res.Body != "{}" {
				t.Fatal("Expected default response body to be empty object")
			}
			return nil
		}
		New(h)(events.APIGatewayProxyRequest{})
	})

	t.Run("should allow cross origin", func(t *testing.T) {
		var h Handler = func(req *Request, res *Response) *Error {
			if res.Headers["Access-Control-Allow-Origin"] != "*" {
				t.Fatal("Expected default content type to be json")
			}
			return nil
		}
		New(h)(events.APIGatewayProxyRequest{})
	})

	t.Run("should transform handler errors into valid responses", func(t *testing.T) {
		message := "error message"
		status := 123
		var h Handler = func(req *Request, res *Response) *Error {
			return &Error{
				code:  status,
				error: errors.New(message),
			}
		}

		res, _ := New(h)(events.APIGatewayProxyRequest{})
		if res.StatusCode != status {
			t.Fatal("Error status was not used in response")
		}
		if res.Body != message {
			t.Fatal("Error message was not used in response")
		}
		if res.Headers["Content-Type"] != "text/plain" {
			t.Fatal("Expected default content type to be plaintext")
		}
	})
}
