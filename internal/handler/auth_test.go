package handler

import (
	"encoding/base64"
	"testing"
	"time"

	mockSecrets "github.com/g-harel/targetblank/services/secrets/mock"
)

func init() {
	longTTL = time.Millisecond * 16
	shortTTL = time.Millisecond * 4
}

func TestCreateToken(t *testing.T) {
	t.Run("should not produce the same token for the same input", func(t *testing.T) {
		secret := "test secret"

		tkn1, err := CreateToken(mockSecrets.RawKey, false, secret)
		if err != nil {
			t.Fatalf("Unexpected error creating token: %v", err)
		}

		tkn2, err := CreateToken(mockSecrets.RawKey, false, secret)
		if err != nil {
			t.Fatalf("Unexpected error creating token: %v", err)
		}

		if tkn1 == tkn2 {
			t.Fatal("Token value should be different")
		}
	})

	t.Run("should produce base64 encoded tokens", func(t *testing.T) {
		secrets := []string{
			"test secret 1",
			"test secret 2",
			"test secret 3",
		}

		for _, secret := range secrets {
			tkn, funcErr := CreateToken(mockSecrets.RawKey, false, secret)
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
	Authenticate := func(token, secret string) *Error {
		req := &Request{
			Headers: map[string]string{},
		}
		req.Headers[AuthHeader] = AuthType + " " + token
		return req.Authenticate(mockSecrets.RawKey, secret)
	}

	t.Run("should produce an error if the secret is wrong", func(t *testing.T) {
		secret := "s3cr3t"

		tkn, err := CreateToken(mockSecrets.RawKey, false, secret)
		if err != nil {
			t.Fatalf("Unexpected error creating token: %v", err)
		}
		funcErr := Authenticate(tkn, secret)
		if funcErr != nil {
			t.Fatalf("Unexpected error when validating with a correct secret: %v", funcErr)
		}

		tkn, err = CreateToken(mockSecrets.RawKey, false, secret)
		if err != nil {
			t.Fatalf("Unexpected error creating token: %v", err)
		}
		funcErr = Authenticate(tkn, "wrong secret")
		if funcErr == nil {
			t.Fatal("Expected incorrect secret to produce error")
		}
	})

	t.Run("should reject expired tokens", func(t *testing.T) {
		secret := "secret"

		tkn, err := CreateToken(mockSecrets.RawKey, false, secret)
		if err != nil {
			t.Fatalf("Unexpected error creating token: %v", err)
		}

		funcErr := Authenticate(tkn, secret)
		if funcErr != nil {
			t.Fatalf("Unexpected error reading token: %v", funcErr)
		}

		time.Sleep(longTTL)

		funcErr = Authenticate(tkn, secret)
		if funcErr == nil {
			t.Fatalf("Expected expired token to be rejected")
		}
	})

	t.Run("should use short expiry for restricted tokens", func(t *testing.T) {
		secret := "SeCrEt"

		tkn, err := CreateToken(mockSecrets.RawKey, true, secret)
		if err != nil {
			t.Fatalf("Unexpected error creating token: %v", err)
		}

		funcErr := Authenticate(tkn, secret)
		if funcErr != nil {
			t.Fatalf("Unexpected error reading token: %v", funcErr)
		}

		time.Sleep(shortTTL)

		funcErr = Authenticate(tkn, secret)
		if funcErr == nil {
			t.Fatalf("Expected expired token to be rejected")
		}
	})
}
