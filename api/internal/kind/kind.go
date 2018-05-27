package kind

import (
	"errors"
	"regexp"
	"strings"
)

type kind int

// Kinds that can be checked.
const (
	EMAIL = kind(iota)
	PASSWORD
)

type pattern struct {
	pattern *regexp.Regexp
	message string
}

var validationKinds = map[kind]*pattern{
	EMAIL: &pattern{
		pattern: regexp.MustCompile(`^\S+@\S+\.\S+$`),
		message: "value does not match an email address (address@domain.tld)",
	},
	PASSWORD: &pattern{
		pattern: regexp.MustCompile(`^.{8,}$`),
		message: "password is shorter than eight characters",
	},
}

type state struct {
	input string
}

func (s *state) Is(k kind) error {
	v := validationKinds[k]
	if v == nil {
		return errors.New("validation kind unknown")
	}
	if !v.pattern.Match([]byte(strings.TrimSpace(s.input))) {
		return errors.New(v.message)
	}
	return nil
}

// Of creates a new checker state.
func Of(s string) *state {
	return &state{s}
}
