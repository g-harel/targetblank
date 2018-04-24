package pageparser

import (
	"github.com/g-harel/targetblank/page"
	"github.com/g-harel/targetblank/parser"
)

// Parser creates a new parser binded to a page object.
func Parser(pg *page.Page) *parser.Parser {
	ps := parser.New()
	ps.Add(
		Empty(pg),
		Whitespace(pg),
		Comment(pg),
		Version(pg),
	)
	return ps
}
