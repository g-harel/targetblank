package parser

import (
	"regexp"
)

// Rule defines behavior associated with patterns in the input string.
type Rule struct {
	Name     string
	Pattern  *regexp.Regexp
	Handler  func(ctx *Context)
	Required bool
	disabled bool // Unexported to give Context exclusive control to change.
}
