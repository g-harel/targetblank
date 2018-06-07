package hash

import "golang.org/x/crypto/bcrypt"

var hashComplexity = 10

// New hashes the input string.
func New(s string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(s), hashComplexity)
	return string(res), err
}

// Check compares the input string to the input hash.
func Check(s, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(s)) == nil
}