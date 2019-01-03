package function

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/g-harel/targetblank/internal/crypto"
)

var headerName = "token"
var longTTL = time.Hour * 18
var shortTTL = time.Minute * 10

type tokenPayload struct {
	ExpireAt   int64  `json:"a"`
	Restricted bool   `json:"b"`
	Secret     string `json:"c"`
}

// MakeToken creates a new authentication token.
func MakeToken(restricted bool, secret string) (string, *Error) {
	expire := time.Now()
	if restricted {
		expire = expire.Add(shortTTL)
	} else {
		expire = expire.Add(longTTL)
	}

	payload, err := json.Marshal(&tokenPayload{
		ExpireAt:   expire.UnixNano(),
		Restricted: restricted,
		Secret:     secret,
	})
	if err != nil {
		return "", Err(http.StatusInternalServerError, err)
	}

	t, err := crypto.Encrypt(payload)
	if err != nil {
		return "", Err(http.StatusInternalServerError, err)
	}

	return t, nil
}

// ValidateToken validates the token in the request.
func (r *Request) ValidateToken(secret string) (restricted bool, e *Error) {
	t := r.Headers[headerName]
	if t == "" {
		return false, CustomErr(errors.New("missing authentication token"))
	}

	payload, err := crypto.Decrypt(t)
	if err != nil {
		return false, Err(http.StatusBadRequest, err)
	}

	p := &tokenPayload{}
	err = json.Unmarshal(payload, p)
	if err != nil {
		return false, Err(http.StatusInternalServerError, err)
	}

	if p.ExpireAt < time.Now().UnixNano() {
		return false, Err(http.StatusBadRequest, errors.New("expired token"))
	}
	if p.Secret != secret {
		return false, Err(http.StatusBadRequest, errors.New("incorrect token secret"))
	}

	return p.Restricted, nil
}
