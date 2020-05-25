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

// Changing the json encoding field names will break issued tokens.
type tokenPayload struct {
	ExpireAt int64  `json:"a"`
	Secret   string `json:"b"`
	IssuedAt int64  `json:"c"`
}

// TODO CreatePasswordToken that can only be used to update the password and has a shorter timeout.

// CreateToken creates a new authentication token.
func CreateToken(key string, short bool, secret string) (string, error) {
	expire := time.Now()
	if short {
		expire = expire.Add(shortTTL)
	} else {
		expire = expire.Add(longTTL)
	}

	payload, err := json.Marshal(&tokenPayload{
		ExpireAt: expire.UnixNano(),
		Secret:   secret,
		IssuedAt: time.Now().UnixNano(),
	})
	if err != nil {
		return "", fmt.Errorf("marshall token: %v", err)
	}

	t, err := crypto.Encrypt(key, payload)
	if err != nil {
		return "", fmt.Errorf("encrypt token: %v", err)
	}

	return t, nil
}

// Authenticate validates the token in the request.
func (r *Request) Authenticate(key, secret string) (*time.Time, *Error) {
	raw := r.Headers[AuthHeader]
	if raw == "" {
		return nil, ClientErr("missing authorization (no \"%v\" header)", AuthHeader)
	}

	values := strings.Fields(raw)
	if len(values) < 2 || values[0] != AuthType {
		return nil, ClientErr("invalid authorization")
	}

	payload, err := crypto.Decrypt(key, values[1])
	if err != nil {
		return nil, ClientErr("invalid authorization")
	}

	token := &tokenPayload{}
	err = json.Unmarshal(payload, token)
	if err != nil {
		return nil, ClientErr("invalid authorization")
	}

	if token.ExpireAt < time.Now().UnixNano() {
		return nil, ClientErr("expired authorization")
	}
	if token.Secret != secret {
		return nil, ClientErr("invalid authorization")
	}

	issuedAt := time.Unix(token.IssuedAt/int64(time.Second), 0)
	return &issuedAt, nil
}
