package parser

import (
	"fmt"
)

type Context struct {
	lines         []string
	parser        *Parser
	ignoredLines  int
	currentRule   *Rule
	currentParams map[string]string
	currentErr    error
}

func (c *Context) reset() {
	c.currentRule = nil
	c.currentErr = nil
	c.currentParams = make(map[string]string)
}

func (c *Context) IgnoreLine() {
	c.lines = c.lines[1:]
	c.ignoredLines++
}

func (c *Context) ReplaceLine(s string) {
	c.lines[0] = s
}

func (c *Context) AddRules(r ...*Rule) {
	c.parser.Add(r...)
}

func (c *Context) DisableRule() {
	c.currentRule.disabled = true
}

func (c *Context) DisableOtherRule(name string) {
	for _, r := range c.parser.rules {
		if r.name == name {
			r.disabled = true
		}
	}
}

func (c *Context) Param(s string) string {
	return c.currentParams[s]
}

func (c *Context) Error(s string, args ...interface{}) {
	errString := fmt.Sprintf(s, args...)
	c.currentErr = fmt.Errorf("Error on line %d: %v", c.ignoredLines+1, errString)
}
