package handler

import (
	"encoding/base64"
	"testing"
	"time"

	mockSecrets "github.com/g-harel/targetblank/services/secrets/mock"
)

func init() {
	restrictedTokenTTL = time.Millisecond * 4
}

func TestCreateToken(t *testing.T) {
	t.Run("should not produce the same token for the same input", func(t *testing.T) {
		identity := "test identity"

		tkn1, err := CreateToken(mockSecrets.RawKey, identity)
		if err != nil {
			t.Fatalf("Unexpected error creating token: %v", err)
		}

		tkn2, err := CreateToken(mockSecrets.RawKey, identity)
		if err != nil {
			t.Fatalf("Unexpected error creating token: %v", err)
		}

		if tkn1 == tkn2 {
			t.Fatal("Token value should be different")
		}
	})

	t.Run("should produce base64 encoded tokens", func(t *testing.T) {
		identities := []string{
			"test identity 1",
			"test identity 2",
			"test identity 3",
		}

		for _, identity := range identities {
			tkn, funcErr := CreateToken(mockSecrets.RawKey, identity)
			if funcErr != nil {
				t.Fatalf("Unexpected error creating token: %v", funcErr)
			}

			_, err := base64.URLEncoding.DecodeString(tkn)
			if err != nil {
				t.Fatalf("Failed to decode base64 token")
			}
		}
	})
}

func TestAuthenticate(t *testing.T) {
	Authenticate := func(token, identity string) error {
		req := &Request{
			Headers: map[string]string{},
		}
		req.Headers[AuthHeader] = AuthType + " " + token
		_, err := req.Authenticate(mockSecrets.RawKey, identity)
		return err
	}

	t.Run("should produce an error if the identity is wrong", func(t *testing.T) {
		identity := "1d3nt1ty"

		tkn, err := CreateToken(mockSecrets.RawKey, identity)
		if err != nil {
			t.Fatalf("Unexpected error creating token: %v", err)
		}
		funcErr := Authenticate(tkn, identity)
		if funcErr != nil {
			t.Fatalf("Unexpected error when validating with a correct identity: %v", funcErr)
		}

		tkn, err = CreateToken(mockSecrets.RawKey, identity)
		if err != nil {
			t.Fatalf("Unexpected error creating token: %v", err)
		}
		funcErr = Authenticate(tkn, "wrong identity")
		if funcErr == nil {
			t.Fatal("Expected incorrect identity to produce error")
		}
	})

	t.Run("should reject expired restricted tokens", func(t *testing.T) {
		identity := "identity"

		tkn, err := CreateRestrictedToken(mockSecrets.RawKey, identity)
		if err != nil {
			t.Fatalf("Unexpected error creating token: %v", err)
		}

		funcErr := Authenticate(tkn, identity)
		if funcErr != ErrRestrictedToken {
			t.Fatalf("Unexpected error reading token: %v", funcErr)
		}

		time.Sleep(restrictedTokenTTL)

		funcErr = Authenticate(tkn, identity)
		if funcErr == ErrRestrictedToken {
			t.Fatalf("Expected expired token to be rejected")
		}
	})
}
