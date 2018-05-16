package function

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/g-harel/targetblank/internal/token"
)

var tokenHeaderName = "Token"

type tokenPayload struct {
	Expiry int `json:"expiry"`
}

// ValidateToken validates the token in the request.
func (r *Request) ValidateToken() *Error {
	t := r.Headers[tokenHeaderName]

	payload, err := token.Open(t)
	if err != nil {
		return Err(http.StatusForbidden, err)
	}

	p := tokenPayload{}
	err = json.Unmarshal(payload, p)
	if err != nil {
		return Err(http.StatusInternalServerError, err)
	}

	if p.Expiry < 0 { // TODO check expiry
		return Err(http.StatusForbidden, errors.New("expired token"))
	}

	return nil
}

// SendToken sends a new authentication token in the response.
func (r *Response) SendToken() *Error {
	payload, err := json.Marshal(&tokenPayload{
		Expiry: 1, // TODO set expiry
	})
	if err != nil {
		return Err(http.StatusInternalServerError, err)
	}

	t, err := token.Seal(payload)
	if err != nil {
		return Err(http.StatusInternalServerError, err)
	}

	r.Headers["Content-Type"] = "text/plain"
	r.Body = t
	return nil
}
