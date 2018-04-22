package rules

import (
	"github.com/g-harel/targetblank/page"
	"github.com/g-harel/targetblank/parser"
)

// Metadata matches with header metadata values.
func Metadata(p *page.Page) *parser.Rule {
	return parser.NewRule().
		Name("metadata").
		Pattern("(?P<key>[A-Za-z0-9_-]+)=\"(?P<value>.*)\"").
		Handler(func(ctx *parser.Context) {
			p.AddMeta(ctx.Param("key"), ctx.Param("value"))
			ctx.IgnoreLine()
		})
}

// Header is a required rule which matches with the header delimiter.
// Once the header is found, the remaining rules are added to the parser.
func Header(p *page.Page) *parser.Rule {
	return parser.NewRule().
		Name("header").
		Required().
		Pattern("===").
		Handler(func(ctx *parser.Context) {
			ctx.IgnoreLine()
			ctx.DisableRule()
			ctx.AddRules(
				Group(p),
				Label(p),
			)
		})
}

// Group matches group delimiters.
// These delimiters indicate a new group should be created.
func Group(p *page.Page) *parser.Rule {
	return parser.NewRule().
		Name("group").
		Pattern("---").
		Handler(func(ctx *parser.Context) {
			p.AddGroup()
			ctx.IgnoreLine()
			ctx.DisableOtherRule("metadata")
		})
}

// Label matches labelled links.
// Item are added to the page at the specified depth.
func Label(p *page.Page) *parser.Rule {
	return parser.NewRule().
		Name("label").
		Pattern("(?P<indent>\\s*)(?P<label>[^\\s\\[].+?)?(?:\\[(?P<link>.*)\\])?").
		Handler(func(ctx *parser.Context) {
			indent := ctx.Param("indent")
			label := ctx.Param("label")
			link := ctx.Param("link")

			if len(indent)%4 != 0 {
				ctx.Error("expected indentation to be in 4 space increments")
				return
			}
			depth := len(indent) / 4

			err := p.AddItem(depth, page.NewItem(link, label))
			if err != nil {
				ctx.Error(err.Error())
				return
			}
			ctx.IgnoreLine()
		})
}
