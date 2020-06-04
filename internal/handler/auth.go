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

var (
	restrictedTokenTTL = time.Minute * 10

	// ErrRestrictedToken is returned when authenticating restricted tokens.
	ErrRestrictedToken = fmt.Errorf("restricted token")
)

// Changing the json encoding field names will break issued tokens.
type tokenPayload struct {
	ExpireAt int64  `json:"a"`
	Identity string `json:"b"`
	IssuedAt int64  `json:"c"`
}

// CreateRestrictedToken creates a new authentication token that expires.
func CreateRestrictedToken(key string, identity string) (string, error) {
	return createToken(key, identity, &restrictedTokenTTL)
}

// CreateToken creates a new authentication token that never expires.
func CreateToken(key string, identity string) (string, error) {
	return createToken(key, identity, nil)
}

func createToken(key string, identity string, duration *time.Duration) (string, error) {
	var expireAt int64
	if duration != nil {
		expireAt = time.Now().Add(*duration).UnixNano()
	}

	payload, err := json.Marshal(&tokenPayload{
		ExpireAt: expireAt,
		Identity: identity,
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
// Returns time at which token was issued and/or an authentication error.
// Returns ErrRestrictedToken if token is restricted.
func (r *Request) Authenticate(key, identity string) (*time.Time, error) {
	raw := r.Headers[AuthHeader]
	if raw == "" {
		return nil, fmt.Errorf("missing authorization (no \"%v\" header)", AuthHeader)
	}

	values := strings.Fields(raw)
	if len(values) < 2 || values[0] != AuthType {
		return nil, fmt.Errorf("could not parse auth header")
	}

	payload, err := crypto.Decrypt(key, values[1])
	if err != nil {
		return nil, fmt.Errorf("invalid authorization")
	}

	token := &tokenPayload{}
	err = json.Unmarshal(payload, token)
	if err != nil {
		return nil, fmt.Errorf("unmarshal: %v", err)
	}

	if token.ExpireAt != 0 && token.ExpireAt < time.Now().UnixNano() {
		return nil, fmt.Errorf("expired token")
	}
	if token.Identity != identity {
		return nil, fmt.Errorf("incorrect identity")
	}

	issuedAt := time.Unix(token.IssuedAt/int64(time.Second), 0)

	if token.ExpireAt != 0 {
		return &issuedAt, ErrRestrictedToken
	}

	return &issuedAt, nil
}
