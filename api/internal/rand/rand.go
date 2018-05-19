package rand

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var randBucket = []rune(
	"0123456789" +
		"abcdefghijkmnopqrstuvwxyz" +
		"ABCDEFGHJKLMNOPQRSTUVWXYZ")

// String generates a pseudorandom string of the specified length.
func String(l int) string {
	b := make([]rune, l)
	for i := range b {
		b[i] = randBucket[rand.Intn(len(randBucket))]
	}
	return string(b)
}
