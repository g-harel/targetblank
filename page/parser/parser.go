package pageparser

import (
	"github.com/g-harel/targetblank/page"
	"github.com/g-harel/targetblank/parser"
)

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
