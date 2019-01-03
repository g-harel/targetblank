package crypto

import (
	"testing"
)

func TestHash(t *testing.T) {
	t.Run("should not produce the same hash for the same input", func(t *testing.T) {
		value := "test value"

		hash1, err := Hash(value)
		if err != nil {
			t.Fatalf("Error creating hash: %v", err)
		}

		hash2, err := Hash(value)
		if err != nil {
			t.Fatalf("Error creating hash: %v", err)
		}

		if hash1 == hash2 {
			t.Fatal("Hash value should be different")
		}
	})
}

func TestHashCheck(t *testing.T) {
	t.Run("should correctly match hashed values", func(t *testing.T) {
		strings := []string{
			"test value",
			"Testing «ταБЬℓσ»: 1<2 & 4+1>3, now 20%% off!",
			"٩(-̮̮̃-̃)۶ ٩(●̮̮̃•̃)۶ ٩(͡๏̯͡๏)۶ ٩(-̮̮̃•̃).",
		}

		for _, s := range strings {
			hash, err := Hash(s)
			if err != nil {
				t.Fatalf("Error hashing string: %v", err)
			}

			match := HashCheck(s, hash)
			if !match {
				t.Fatal("Matching hash/value pair was not successfully checked")
			}
		}
	})

	t.Run("should correctly reject incorrect values", func(t *testing.T) {
		strings := []string{
			"test value 1",
			"test value 2",
			"test value 3",
		}

		hash, err := Hash("vest talue")
		if err != nil {
			t.Fatalf("Error hashing string: %v", err)
		}

		for _, s := range strings {
			match := HashCheck(s, hash)
			if match {
				t.Fatal("Incorrect hash/value pair was not rejected")
			}
		}
	})
}
