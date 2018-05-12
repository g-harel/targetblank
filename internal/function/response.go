package function

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// Response replaces the api gateway response event.
type Response events.APIGatewayProxyResponse

// ClientErr adds a client error to the response.
func (r *Response) ClientErr(status int, err error) {
	fmt.Println(err) // TODO remove
	r.StatusCode = status
	if r.Headers == nil {
		r.Headers = map[string]string{}
	}
	r.Headers["Content-Type"] = "text/plain"
	r.Body = err.Error()
}

// ServerErr adds a server error to the response.
func (r *Response) ServerErr(status int, err error) {
	fmt.Println(err) // TODO proper logging
	r.StatusCode = status
	if r.Headers == nil {
		r.Headers = map[string]string{}
	}
	r.Headers["Content-Type"] = "text/plain"
	r.Body = http.StatusText(status)
}
