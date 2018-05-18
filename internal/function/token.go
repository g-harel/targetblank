package function

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/g-harel/targetblank/internal/token"
)

var headerName = "Token"
var longTTL = time.Hour * 24
var shortTTL = time.Minute * 5

type tokenPayload struct {
	ExpireAt   int64  `json:"a"`
	Restricted bool   `json:"b"`
	Secret     string `json:"c"`
}

// MakeToken creates a new authentication token.
func MakeToken(restricted bool, secret string) (string, *Error) {
	expire := time.Now()
	if restricted {
		expire.Add(shortTTL)
	} else {
		expire.Add(longTTL)
	}

	payload, err := json.Marshal(&tokenPayload{
		ExpireAt:   expire.Unix(),
		Restricted: restricted,
		Secret:     secret,
	})
	if err != nil {
		return "", Err(http.StatusInternalServerError, err)
	}

	t, err := token.Seal(payload)
	if err != nil {
		return "", Err(http.StatusInternalServerError, err)
	}

	return t, nil
}

// ValidateToken validates the token in the request.
// Token validity errors are always 404 to avoid leaking the existence of unpublished pages.
func (r *Request) ValidateToken(secret string) (restricted bool, e *Error) {
	t := r.Headers[headerName]
	if t == "" {
		return false, CustomErr(errors.New("missing authentication token"))
	}

	payload, err := token.Open(t)
	if err != nil {
		return false, Err(http.StatusNotFound, err)
	}

	p := &tokenPayload{}
	err = json.Unmarshal(payload, p)
	if err != nil {
		return false, Err(http.StatusInternalServerError, err)
	}

	if p.ExpireAt < time.Now().Unix() {
		return false, Err(http.StatusNotFound, errors.New("expired token"))
	}
	if p.Secret != secret {
		return false, Err(http.StatusNotFound, errors.New("incorrect token secret"))
	}

	return p.Restricted, nil
}
