package parser

import (
	"fmt"
	"strings"
)

type Parser struct {
	rules []*Rule
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) Add(rules ...*Rule) {
	p.rules = append(p.rules, rules...)
}

func (p *Parser) Parse(s string) error {
	lines := strings.Split(s, "\n")
	if len(lines) == 0 {
		return fmt.Errorf("input string is empty")
	}

	ctx := &Context{
		lines:  lines,
		parser: p,
	}

	for len(ctx.lines) > 0 {
		matched := false
		for _, r := range p.rules {
			if r.disabled || r.pattern == nil || r.handler == nil {
				continue
			}
			match := r.pattern.FindStringSubmatch(ctx.lines[0])
			if match == nil {
				if r.required {
					ctx.Error("expected %s", r.name)
					return ctx.currentErr
				}
				continue
			}
			matched = true

			ctx.reset()
			ctx.currentRule = r
			for i, s := range r.pattern.SubexpNames() {
				if i <= len(match) {
					ctx.currentParams[s] = match[i]
				}
			}

			(*r.handler)(ctx)
			if ctx.currentErr != nil {
				return ctx.currentErr
			}
			break
		}
		if !matched {
			ctx.Error("syntax error")
			return ctx.currentErr
		}
	}

	for _, r := range p.rules {
		if r.required && !r.disabled {
			ctx.Error("expected %s", r.name)
			return ctx.currentErr
		}
	}

	return nil
}
