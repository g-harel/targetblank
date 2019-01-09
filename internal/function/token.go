package function

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/g-harel/targetblank/internal/crypto"
)

const tokenHeader = "token"

var longTTL = time.Hour * 18
var shortTTL = time.Minute * 10

type tokenPayload struct {
	ExpireAt int64  `json:"a"`
	Secret   string `json:"b"`
}

// CreateToken creates a new authentication token.
func CreateToken(short bool, secret string) (string, error) {
	expire := time.Now()
	if short {
		expire = expire.Add(shortTTL)
	} else {
		expire = expire.Add(longTTL)
	}

	payload, err := json.Marshal(&tokenPayload{
		ExpireAt: expire.UnixNano(),
		Secret:   secret,
	})
	if err != nil {
		return "", fmt.Errorf("marshall token: %v", err)
	}

	t, err := crypto.Encrypt(payload)
	if err != nil {
		return "", fmt.Errorf("encrypt token: %v", err)
	}

	return t, nil
}

// Authenticate validates the token in the request.
func (r *Request) Authenticate(secret string) *Error {
	t := r.Headers[tokenHeader]
	if t == "" {
		return ClientErr("missing authentication token (no \"%v\" header)", tokenHeader)
	}

	payload, err := crypto.Decrypt(t)
	if err != nil {
		return ClientErr("invalid authentication token")
	}

	p := &tokenPayload{}
	err = json.Unmarshal(payload, p)
	if err != nil {
		return ClientErr("invalid authentication token")
	}

	if p.ExpireAt < time.Now().UnixNano() {
		return ClientErr("expired authentication token")
	}
	if p.Secret != secret {
		return ClientErr("invalid authentication token")
	}

	return nil
}
