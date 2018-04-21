package parser

import "regexp"

type RuleHandler func(ctx *Context)

type rule struct {
	pattern       *regexp.Regexp
	handler       *RuleHandler
	disabled      bool
	strict        bool
	strictMessage string
}

func NewRule(pattern string, handler RuleHandler) *rule {
	return (&rule{}).Pattern(pattern).Handler(handler)
}

func (r *rule) Pattern(p string) *rule {
	r.pattern = regexp.MustCompile(p)
	return r
}

func (r *rule) Handler(h RuleHandler) *rule {
	r.handler = &h
	return r
}

func (r *rule) Strict(m string) *rule {
	r.strict = true
	r.strictMessage = m
	return r
}

func (r *rule) Disable() *rule {
	r.disabled = true
	return r
}
