package rand

import (
	"math/rand"
	"regexp"
	"testing"
)

func TestString(t *testing.T) {
	t.Run("should produce strings of the requested length", func(t *testing.T) {
		for i := 0; i < 32; i++ {
			l := rand.Intn(128)
			v := String(l)
			if len(v) != l {
				t.Fatal("Produced string does not have the correct length")
			}
		}
	})

	t.Run("should produce random strings", func(t *testing.T) {
		dict := map[string]bool{}

		for i := 0; i < 32; i++ {
			v := String(16)
			if dict[v] == true {
				t.Fatal("Duplicate value produced")
			}
			dict[v] = true
		}
	})

	t.Run("should produce only letters and numbers", func(t *testing.T) {
		p := regexp.MustCompile(`[0-9a-zA-Z]+`)

		for i := 0; i < 32; i++ {
			v := String(16)
			if !p.Match([]byte(v)) {
				t.Fatal("Duplicate value produced")
			}
		}
	})
}
