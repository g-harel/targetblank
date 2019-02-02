package handler

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/g-harel/targetblank/internal/crypto"
)

// Expected authorization header configuration.
const (
	AuthHeader = "Authorization"
	AuthType   = "Targetblank"
)

var longTTL = time.Hour * 24 * 3
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
	raw := r.Headers[AuthHeader]
	if raw == "" {
		return ClientErr("missing authorization (no \"%v\" header)", AuthHeader)
	}

	values := strings.Fields(raw)
	if len(values) < 2 || values[0] != AuthType {
		return ClientErr("invalid authorization")
	}

	payload, err := crypto.Decrypt(values[1])
	if err != nil {
		return ClientErr("invalid authorization")
	}

	token := &tokenPayload{}
	err = json.Unmarshal(payload, token)
	if err != nil {
		return ClientErr("invalid authorization")
	}

	if token.ExpireAt < time.Now().UnixNano() {
		return ClientErr("expired authorization")
	}
	if token.Secret != secret {
		return ClientErr("invalid authorization")
	}

	return nil
}
