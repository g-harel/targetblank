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

func (p *parser) Add(r *rule) {
	p.rules = append(p.rules, r)
}

func (p *parser) Parse(s string) error {
	lines := strings.Split(s, "\n")
	if len(lines) == 0 {
		return fmt.Errorf("input string is empty")
	}

	ctx := &Context{
		lineNumber: 0,
		lines:      lines,
	}

	for ctx.lineNumber >= 0 && ctx.lineNumber < len(ctx.lines) {
		matched := false
		for _, r := range p.rules {
			if r.disabled || r.pattern == nil || r.handler == nil {
				continue
			}
			match := r.pattern.FindStringSubmatch(ctx.lines[ctx.lineNumber])
			if match == nil {
				if r.strict {
					return fmt.Errorf(r.strictMessage)
				}
				continue
			}
			matched = true

			ctx.params = make(map[string]string)
			for i, s := range r.pattern.SubexpNames() {
				if i <= len(match) {
					ctx.params[s] = match[i]
				}
			}
			ctx.currentRule = r
			(*r.handler)(ctx)
			if ctx.err != nil {
				return ctx.err
			}
			break
		}
		if !matched {
			return fmt.Errorf("could not match line: %s", ctx.lines[ctx.lineNumber])
		}
	}

	return nil
}
