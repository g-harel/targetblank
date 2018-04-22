package parser

import (
	"fmt"
	"strings"
)

type parser struct {
	rules []*rule
}

func New() *parser {
	return &parser{}
}

func (p *parser) Add(rules ...*rule) {
	p.rules = append(p.rules, rules...)
}

func (p *parser) Parse(s string) error {
	lines := strings.Split(s, "\n")
	if len(lines) == 0 {
		return fmt.Errorf("input string is empty")
	}

	ctx := &Context{
		lines: lines,
	}

	for len(ctx.lines) > 0 {
		matched := false
		for _, r := range p.rules {
			if r.disabled || r.pattern == nil || r.handler == nil {
				continue
			}
			match := r.pattern.FindStringSubmatch(ctx.lines[0])
			if match == nil {
				if r.strict {
					ctx.Error(r.strictMessage)
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
			ctx.Error("could not match line")
			return ctx.currentErr
		}
	}

	return nil
}
