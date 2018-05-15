package token

import "github.com/g-harel/targetblank/internal/rand"

// Generate creates a new token from the address.
func Generate(addr string) string {
	// TODO
	return rand.String(16)
}

// Validate checks that the token is legitimate.
func Validate(addr string) bool {
	// TODO
	return true
}
