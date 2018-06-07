package parser

import (
	"fmt"
)

// Context allows rule handlers to modify parser state in controlled ways.
type Context struct {
	lines         []string
	parser        *Parser
	ignoredLines  int
	currentRule   *Rule
	currentParams map[string]string
	currentErr    *Error
}

// Resets all temporary context state.
func (c *Context) reset() {
	c.currentRule = nil
	c.currentErr = nil
	c.currentParams = make(map[string]string)
}

// LineParsed signals that the line has been processed and can now be ignored.
func (c *Context) LineParsed() {
	c.lines = c.lines[1:]
	c.ignoredLines++
}

// ReplaceLine changes the content of the matched line.
// The parser will re-check the modified line on the rules.
func (c *Context) ReplaceLine(s string) {
	c.lines[0] = s
}

// DisableSelf disables the rule associated with the matcher calling it.
func (c *Context) DisableSelf() {
	c.currentRule.Disabled = true
}

// DisableOther disables rules on the parent parser.
// This can be used to disallow rules after a certain marker.
func (c *Context) DisableOther(name string) {
	for _, r := range c.parser.rules {
		if r.Name == name {
			r.Disabled = true
		}
	}
}

// EnableOther allows rule handlers to enable more rules once they've matched.
// This functionality (coupled with rule disabling) makes it easier to parse different sections of files independently.
func (c *Context) EnableOther(name string) {
	for _, r := range c.parser.rules {
		if r.Name == name {
			r.Disabled = false
		}
	}
}

// Param fetches the value from the named capture group in the rule's pattern.
func (c *Context) Param(s string) string {
	return c.currentParams[s]
}

// Error adds an error to the context and formats it to include the line number.
func (c *Context) Error(s string, args ...interface{}) {
	c.currentErr = &Error{
		Line:  c.ignoredLines + 1,
		error: fmt.Errorf(s, args...),
	}
}