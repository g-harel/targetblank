package parser

import (
	"fmt"
)

type Context struct {
	lines         []string
	ignoredLines  int
	currentRule   *rule
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

func (c *Context) DisableRule() {
	c.currentRule.Disable()
}

func (c *Context) Param(s string) string {
	return c.currentParams[s]
}

func (c *Context) Error(s string, args ...interface{}) {
	errString := fmt.Sprintf(s, args...)
	c.currentErr = fmt.Errorf("Error on line %d: %v", c.ignoredLines+1, errString)
}
