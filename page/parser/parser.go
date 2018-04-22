package pageparser

import (
	"github.com/g-harel/targetblank/page"
	"github.com/g-harel/targetblank/parser"
)

// Parse uses this package's rules to parse the input string into a page.
func Parse(s string) (*page.Page, error) {
	pg := page.New()

	ps := parser.New()
	ps.Add(
		Empty(pg),
		Whitespace(pg),
		Comment(pg),
		Version(pg),
	)

	return pg, ps.Parse(s)
}
