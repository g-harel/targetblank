package parser

import "regexp"

type RuleHandler func(ctx *Context, params map[string]string) error

type Rule struct {
	pattern *regexp.Regexp
	handler *RuleHandler
}

func NewRule(pattern string, handler RuleHandler) *Rule {
	return (&Rule{}).Pattern(pattern).Handler(handler)
}

func (r *Rule) Pattern(p string) *Rule {
	r.pattern = regexp.MustCompile(p)
	return r
}

func (r *Rule) Handler(h RuleHandler) *Rule {
	r.handler = &h
	return r
}

func (r *Rule) Disable() *Rule {
	r.pattern = nil
	return r
}
