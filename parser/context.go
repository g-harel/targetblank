package parser

import (
	"fmt"
)

type Context struct {
	lines       []string
	lineNumber  int
	currentRule *rule
	params      map[string]string
	err         error
}

func (c *Context) RemoveLine() {
	c.lines = append(
		c.lines[:c.lineNumber],
		c.lines[1+c.lineNumber:]...,
	)
}

func (c *Context) ReplaceLine(s string) {
	c.lines[c.lineNumber] = s
}

func (c *Context) DisableRule() {
	c.currentRule.Disable()
}

func (c *Context) Param(s string) string {
	return c.params[s]
}

func (c *Context) Error(s string, args ...interface{}) {
	c.err = fmt.Errorf(s, args...)
}
