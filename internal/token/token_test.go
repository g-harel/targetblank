package token

import (
	"encoding/base64"
	"testing"
)

func TestSeal(t *testing.T) {
	t.Run("should not produce the same token for the same input", func(t *testing.T) {
		payload := "test payload"

		tkn1, err := Seal([]byte(payload))
		if err != nil {
			t.Fatalf("Error creating token: %v", err)
		}

		tkn2, err := Seal([]byte(payload))
		if err != nil {
			t.Fatalf("Error creating token: %v", err)
		}

		if tkn1 == tkn2 {
			t.Fatal("Token value should be different")
		}
	})

	t.Run("should produce base64 encoded tokens", func(t *testing.T) {
		payloads := []string{
			"test payload 1",
			"test payload 2",
			"test payload 3",
		}

		for _, payload := range payloads {
			tkn, err := Seal([]byte(payload))
			if err != nil {
				t.Fatalf("Error creating token: %v", err)
			}

			_, err = base64.URLEncoding.DecodeString(tkn)
			if err != nil {
				t.Fatalf("Failed to decode base64 token")
			}
		}
	})
}

func TestOpen(t *testing.T) {
	t.Run("should correctly decrypt encrypted payloads", func(t *testing.T) {
		payloads := []string{
			"test payload",
			"Testing «ταБЬℓσ»: 1<2 & 4+1>3, now 20%% off!",
			"٩(-̮̮̃-̃)۶ ٩(●̮̮̃•̃)۶ ٩(͡๏̯͡๏)۶ ٩(-̮̮̃•̃).",
		}

		for _, payload := range payloads {
			tkn, err := Seal([]byte(payload))
			if err != nil {
				t.Fatalf("Error creating token: %v", err)
			}

			pld, err := Open(tkn)
			if err != nil {
				t.Fatalf("Error reading token: %v", err)
			}

			if string(pld) != payload {
				t.Fatal("Token value does not match")
			}
		}
	})
}
