package database

import (
	"errors"
	"regexp"
	"strings"
)

type validation struct {
	pattern *regexp.Regexp
	message string
}

var validationKinds = map[string]*validation{
	"email": &validation{
		pattern: regexp.MustCompile("^\\S+@\\S+\\.\\S+$"),
		message: "value does not match an email address (address@domain.tld)",
	},
	"password": &validation{
		pattern: regexp.MustCompile("^.{8,}$"),
		message: "password is shorter than eight characters",
	},
}

// Validate checks that the input string has the correct format for it's decalred kind.
func Validate(s, kind string) error {
	v := validationKinds[kind]
	if v == nil {
		return errors.New("validation kind unknown")
	}
	if !v.pattern.Match([]byte(strings.TrimSpace(s))) {
		return errors.New(v.message)
	}
	return nil
}
