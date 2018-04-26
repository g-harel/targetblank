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
// At parse time, rules which were added the earliest are given the highest priority.
// Added rules are passed by value in order to isolate the original rules, but are stored as references to enable Context to make changes.
func (p *Parser) Add(rules ...Rule) {
	r := make([]*Rule, len(rules))
	for i := range rules {
		r[i] = &rules[i]
	}
	p.rules = append(p.rules, r...)
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

			match := r.Pattern.FindStringSubmatch(ctx.lines[0])
			if match == nil {
				// Required rules must match when checked (rules added before the required ones are still given priority).
				if r.Required {
					ctx.Error("expected %s", r.Name)
					return ctx.currentErr
				}
				continue
			}
			matched = true

			// Context is prepared to be passed to the handler.
			ctx.reset()
			ctx.currentRule = r
			for i, s := range r.Pattern.SubexpNames() {
				if i <= len(match) {
					ctx.currentParams[s] = match[i]
				}
			}

			r.Handler(ctx)
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
		if r.Required && !r.disabled {
			ctx.Error("expected %s", r.Name)
			return ctx.currentErr
		}
	}

	return nil
}
