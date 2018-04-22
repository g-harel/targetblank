package parser

import "regexp"

type RuleHandler func(ctx *Context)

type Rule struct {
	pattern  *regexp.Regexp
	handler  *RuleHandler
	disabled bool
	required bool
	name     string
}

func NewRule() *Rule {
	return &Rule{}
}

func (r *Rule) Name(n string) *Rule {
	r.name = n
	return r
}

func (r *Rule) Pattern(p string) *Rule {
	r.pattern = regexp.MustCompile("^" + p + "$")
	return r
}

func (r *Rule) Handler(h RuleHandler) *Rule {
	r.handler = &h
	return r
}

func (r *Rule) Required() *Rule {
	r.required = true
	return r
}
