package parse

import (
	"fmt"
	"regexp"
	"strings"
)

// Rule defines behavior associated with patterns in the input string.
type rule struct {
	Name     string
	Pattern  *regexp.Regexp
	Handler  func(ctx *context)
	Required bool
	Disabled bool
}

// Parser is used to accumulate rules and parse strings.
type parser struct {
	rules []*rule
}

// Add allows one or more rules to be registered with the parser.
// At parse time, rules which were added the earliest are given the highest priority.
// Added rules are passed by value to isolate the originals, but are stored as references for context's changes to be sticky.
func (p *parser) Add(rules ...rule) *parser {
	r := make([]*rule, len(rules))
	for i := range rules {
		r[i] = &rules[i]
	}
	p.rules = append(p.rules, r...)
	return p
}

// Parse parses the given string according to the added rules.
func (p *parser) Parse(s string) error {
	ctx := &context{
		lines:  strings.Split(s, "\n"),
		parser: p,
	}

	for len(ctx.lines) > 0 {
		matched := false
		for _, r := range p.rules {
			if r.Disabled {
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

		// No match has occurred with any rule's pattern. Line has incorrect syntax.
		if !matched {
			ctx.Error("syntax error")
			return ctx.currentErr
		}
	}

	// If any required rules are still enabled after parsing the full string, content is missing.
	for _, r := range p.rules {
		if r.Required && !r.Disabled {
			ctx.Error("expected %s", r.Name)
			return ctx.currentErr
		}
	}

	return nil
}

// Context allows rule handlers to modify parser state in controlled ways.
type context struct {
	lines         []string
	parser        *parser
	ignoredLines  int
	currentRule   *rule
	currentParams map[string]string
	currentErr    error
}

// Resets all temporary context state.
func (c *context) reset() {
	c.currentRule = nil
	c.currentErr = nil
	c.currentParams = make(map[string]string)
}

// LineParsed signals that the line has been processed and can now be ignored.
func (c *context) LineParsed() {
	c.lines = c.lines[1:]
	c.ignoredLines++
}

// ReplaceLine changes the content of the matched line.
// The parser will re-check the modified line on the rules.
func (c *context) ReplaceLine(s string) {
	c.lines[0] = s
}

// DisableSelf disables the rule associated with the matcher calling it.
func (c *context) DisableSelf() {
	c.currentRule.Disabled = true
}

// DisableOther disables rules on the parent parser.
// This can be used to disallow rules after a certain marker.
func (c *context) DisableOther(name string) {
	for _, r := range c.parser.rules {
		if r.Name == name {
			r.Disabled = true
		}
	}
}

// EnableOther allows rule handlers to enable more rules once they've matched.
// This functionality (coupled with rule disabling) makes it easier to parse different sections of files independently.
func (c *context) EnableOther(name string) {
	for _, r := range c.parser.rules {
		if r.Name == name {
			r.Disabled = false
		}
	}
}

// Param fetches the value from the named capture group in the rule's pattern.
func (c *context) Param(s string) string {
	return c.currentParams[s]
}

// Error adds an error to the context and formats it to include the line number.
func (c *context) Error(s string, args ...interface{}) {
	c.currentErr = fmt.Errorf(
		"line %v: %v",
		c.ignoredLines+1,
		fmt.Sprintf(s, args...),
	)
}
