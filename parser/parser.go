package parser

import (
	"fmt"
	"strings"
)

type parser struct {
	rules []*Rule
}

func New() *parser {
	return &parser{}
}

func (p *parser) Add(r *Rule) {
	p.rules = append(p.rules, r)
}

func (p *parser) AddNewRule(pattern string, handler RuleHandler) {
	p.rules = append(p.rules, NewRule(pattern, handler))
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
			if r.pattern == nil {
				continue
			}
			match := r.pattern.FindStringSubmatch(ctx.lines[ctx.lineNumber])
			if match == nil {
				continue
			}
			matched = true

			params := make(map[string]string)
			for i, s := range r.pattern.SubexpNames() {
				if i <= len(match) {
					params[s] = match[i]
				}
			}
			err := (*r.handler)(ctx, params)
			if err != nil {
				return err
			}
			break
		}
		if !matched {
			return fmt.Errorf("could not match line: %s", ctx.lines[ctx.lineNumber])
		}
	}

	return nil
}
