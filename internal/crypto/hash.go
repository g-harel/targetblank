package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

var cost = 10

// Hash hashes the input string.
func Hash(s string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(s), cost)
	return string(res), err
}

// HashCheck compares the input string to the input hash.
func HashCheck(s, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(s)) == nil
}
