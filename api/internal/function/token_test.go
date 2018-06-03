package function

import (
	"encoding/base64"
	"testing"
	"time"
)

func init() {
	longTTL = time.Millisecond * 3
	shortTTL = time.Millisecond
}

func TestMakeToken(t *testing.T) {
	t.Run("should not produce the same token for the same input", func(t *testing.T) {
		secret := "test secret"

		tkn1, err := MakeToken(false, secret)
		if err != nil {
			t.Fatalf("Unexpected error creating token: %v", err)
		}

		tkn2, err := MakeToken(false, secret)
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
			tkn, funcErr := MakeToken(false, secret)
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

func TestValidateToken(t *testing.T) {
	ValidateToken := func(token, secret string) (bool, *Error) {
		return (&Request{
			Headers: map[string]string{
				headerName: token,
			},
		}).ValidateToken(secret)
	}

	t.Run("should produce an error if the secret is wrong", func(t *testing.T) {
		secret := "s3cr3t"

		tkn, err := MakeToken(false, secret)
		if err != nil {
			t.Fatalf("Unexpected error creating token: %v", err)
		}
		_, err = ValidateToken(tkn, secret)
		if err != nil {
			t.Fatalf("Unexpected error when validating with a correct secret: %v", err)
		}

		tkn, err = MakeToken(false, secret)
		if err != nil {
			t.Fatalf("Unexpected error creating token: %v", err)
		}
		_, err = ValidateToken(tkn, "wrong secret")
		if err == nil {
			t.Fatal("Expected incorrect secret to produce error")
		}
	})

	t.Run("should return the correct restricted status", func(t *testing.T) {
		secret := "secret"

		tkn, err := MakeToken(false, secret)
		if err != nil {
			t.Fatalf("Unexpected error creating token: %v", err)
		}
		r, err := ValidateToken(tkn, secret)
		if err != nil {
			t.Fatalf("Unexpected error reading token: %v", err)
		}
		if r {
			t.Fatalf("Incorrect token status: %v", err)
		}

		tkn, err = MakeToken(true, secret)
		if err != nil {
			t.Fatalf("Unexpected error creating token: %v", err)
		}
		r, err = ValidateToken(tkn, secret)
		if err != nil {
			t.Fatalf("Unexpected error reading token: %v", err)
		}
		if !r {
			t.Fatalf("Incorrect token status: %v", err)
		}
	})

	t.Run("should reject expired tokens", func(t *testing.T) {
		secret := "secret"

		tkn, err := MakeToken(false, secret)
		if err != nil {
			t.Fatalf("Unexpected error creating token: %v", err)
		}

		_, err = ValidateToken(tkn, secret)
		if err != nil {
			t.Fatalf("Unexpected error reading token: %v", err)
		}

		time.Sleep(longTTL)

		_, err = ValidateToken(tkn, secret)
		if err == nil {
			t.Fatalf("Expected expired token to be rejected")
		}
	})

	t.Run("should use short expiry for restricted tokens", func(t *testing.T) {
		secret := "secret"

		tkn, err := MakeToken(true, secret)
		if err != nil {
			t.Fatalf("Unexpected error creating token: %v", err)
		}

		_, err = ValidateToken(tkn, secret)
		if err != nil {
			t.Fatalf("Unexpected error reading token: %v", err)
		}

		time.Sleep(shortTTL)

		_, err = ValidateToken(tkn, secret)
		if err == nil {
			t.Fatalf("Expected expired token to be rejected")
		}
	})
}
