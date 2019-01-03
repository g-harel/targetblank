package rand

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// List of unambiguous characters (minus "Il0O") to use in random strings.
var charBucket = []rune(
	"123456789" +
		"abcdefghijkmnopqrstuvwxyz" +
		"ABCDEFGHJKLMNPQRSTUVWXYZ")

// String generates a pseudorandom string of the specified length.
func String(l int) string {
	b := make([]rune, l)
	for i := range b {
		b[i] = charBucket[rand.Intn(len(charBucket))]
	}
	return string(b)
}
