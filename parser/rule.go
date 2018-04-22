package parser

import "regexp"

// RuleHandler accepts a parser context.
type RuleHandler func(ctx *Context)

// Rule defines behavior associated with patterns in the input string.
type Rule struct {
	name     string
	pattern  *regexp.Regexp
	handler  *RuleHandler
	disabled bool
	required bool
}

// NewRule creates an empty rule object.
func NewRule() *Rule {
	return &Rule{}
}

// Name changes the rule's name.
func (r *Rule) Name(n string) *Rule {
	r.name = n
	return r
}

// Pattern changes the rule's pattern.
// This function also adds tokens to ensure the full line is matched
// The resulting pattern stays valid and correct even if it already contains these tokens.
func (r *Rule) Pattern(p string) *Rule {
	r.pattern = regexp.MustCompile("^" + p + "$")
	return r
}

// Handler changes the rule's handler function.
func (r *Rule) Handler(h RuleHandler) *Rule {
	r.handler = &h
	return r
}

// Required marks the rule as being required.
func (r *Rule) Required() *Rule {
	r.required = true
	return r
}
