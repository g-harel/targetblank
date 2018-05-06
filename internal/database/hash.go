package database

import "golang.org/x/crypto/bcrypt"

var hashComplexity = 12

// Hash hashes the input string.
func Hash(s string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(s), hashComplexity)
	return string(res), err
}

// HashCheck compares the input string to the input hash.
func HashCheck(s, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(s))
	return err == nil
}
