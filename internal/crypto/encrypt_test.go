package crypto

import (
	"encoding/base64"
	"testing"
)

func TestEncrypt(t *testing.T) {
	t.Run("should not produce the same result for the same input", func(t *testing.T) {
		payload := "test payload"

		tkn1, err := Encrypt([]byte(payload))
		if err != nil {
			t.Fatalf("Error creating ciphertext: %v", err)
		}

		tkn2, err := Encrypt([]byte(payload))
		if err != nil {
			t.Fatalf("Error creating ciphertext: %v", err)
		}

		if tkn1 == tkn2 {
			t.Fatal("ciphertext value should be different")
		}
	})

	t.Run("should produce base64 encoded ciphertext", func(t *testing.T) {
		payloads := []string{
			"test payload 1",
			"test payload 2",
			"test payload 3",
		}

		for _, payload := range payloads {
			tkn, err := Encrypt([]byte(payload))
			if err != nil {
				t.Fatalf("Error creating ciphertext: %v", err)
			}

			_, err = base64.URLEncoding.DecodeString(tkn)
			if err != nil {
				t.Fatalf("Failed to decode base64 ciphertext")
			}
		}
	})
}

func TestDecrypt(t *testing.T) {
	t.Run("should correctly decrypt encrypted payloads", func(t *testing.T) {
		payloads := []string{
			"test payload",
			"Testing «ταБЬℓσ»: 1<2 & 4+1>3, now 20%% off!",
			"٩(-̮̮̃-̃)۶ ٩(●̮̮̃•̃)۶ ٩(͡๏̯͡๏)۶ ٩(-̮̮̃•̃).",
		}

		for _, payload := range payloads {
			tkn, err := Encrypt([]byte(payload))
			if err != nil {
				t.Fatalf("Error creating ciphertext: %v", err)
			}

			pld, err := Decrypt(tkn)
			if err != nil {
				t.Fatalf("Error reading ciphertext: %v", err)
			}

			if string(pld) != payload {
				t.Fatal("Ciphertext value does not match")
			}
		}
	})
}
