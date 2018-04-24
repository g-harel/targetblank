package parser

import (
	"strings"
)

// Parser is used to accumulate rules and parse strings.
type Parser struct {
	rules []*Rule
}

// New creates an empty parser.
func New() *Parser {
	return &Parser{}
}

// Add allows one or more rules to be registered with the parser.
// At parse time, rules which are added the earliest are given the highest priority.
func (p *Parser) Add(rules ...*Rule) {
	p.rules = append(p.rules, rules...)
}

// Parse parses the given string according to the added rules.
func (p *Parser) Parse(s string) *Error {
	ctx := &Context{
		lines:  strings.Split(s, "\n"),
		parser: p,
	}

	for len(ctx.lines) > 0 {
		matched := false
		for _, r := range p.rules {
			if r.disabled {
				continue
			}
			if r.pattern == nil || r.handler == nil {
				ctx.Error("invalid rule (name: \"%s\")", r.name)
				return ctx.currentErr
			}

			match := r.pattern.FindStringSubmatch(ctx.lines[0])
			if match == nil {
				// Required rules must match when checked (rules added before the required ones are still given priority).
				if r.required {
					ctx.Error("expected %s", r.name)
					return ctx.currentErr
				}
				continue
			}
			matched = true

			// Context is prepared to be passed to the handler.
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

		// No match has occurred with all rule patterns. This means the line has incorrect syntax.
		if !matched {
			ctx.Error("syntax error")
			return ctx.currentErr
		}
	}

	// If any required rules are still enabled after parsing the full string, content is missing.
	for _, r := range p.rules {
		if r.required && !r.disabled {
			ctx.Error("expected %s", r.name)
			return ctx.currentErr
		}
	}

	return nil
}
