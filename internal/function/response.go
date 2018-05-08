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
	r.StatusCode = status
	r.Body = err.Error() // TODO change Content-Type
}

// ServerErr adds a server error to the response.
func (r *Response) ServerErr(status int, err error) {
	fmt.Println(err) // TODO proper logging
	r.StatusCode = status
	r.Body = http.StatusText(status) // TODO change Content-Type
}
