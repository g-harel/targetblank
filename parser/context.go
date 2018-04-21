package parser

type Context struct {
	lineNumber int
	lines      []string
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
