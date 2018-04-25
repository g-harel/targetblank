package v1

import (
	"github.com/g-harel/targetblank/page"
	"github.com/g-harel/targetblank/parser"
)

// MetadataRule matches with header metadata values.
func MetadataRule(p *page.Page) *parser.Rule {
	return parser.NewRule().
		Name("v1metadata").
		Pattern("(?P<key>[A-Za-z0-9_-]+)\\s*=\\s*(?P<value>.*)\\s*").
		Handler(func(ctx *parser.Context) {
			p.AddMeta(ctx.Param("key"), ctx.Param("value"))
			ctx.IgnoreLine()
		})
}

// HeaderRule is a required rule which matches with the header delimiter.
// Once the header is found, the remaining rules are added to the parser.
func HeaderRule(p *page.Page) *parser.Rule {
	return parser.NewRule().
		Name("v1header").
		Required().
		Pattern("===").
		Handler(func(ctx *parser.Context) {
			ctx.IgnoreLine()
			ctx.DisableRule()
			ctx.AddRules(
				GroupRule(p),
				LabelRule(p),
			)
		})
}

// GroupRule matches group delimiters.
// These delimiters indicate a new group should be created.
func GroupRule(p *page.Page) *parser.Rule {
	return parser.NewRule().
		Name("v1group").
		Pattern("---").
		Handler(func(ctx *parser.Context) {
			p.AddGroup()
			ctx.IgnoreLine()
			ctx.DisableOtherRule("v1metadata")
		})
}

// LabelRule matches labelled links.
// Item are added to the page at the specified depth.
func LabelRule(p *page.Page) *parser.Rule {
	return parser.NewRule().
		Name("v1label").
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
