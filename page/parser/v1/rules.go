package rules

import (
	"strings"

	"github.com/g-harel/targetblank/page"
	"github.com/g-harel/targetblank/parser"
)

func Metadata(p *page.Page) *parser.Rule {
	return parser.NewRule().
		Name("metadata").
		Pattern("(?P<key>[A-Za-z0-9_-]+)=\"(?P<value>.*)\"").
		Handler(func(ctx *parser.Context) {
			p.AddMeta(ctx.Param("key"), ctx.Param("value"))
			ctx.IgnoreLine()
		})
}

func Header(p *page.Page) *parser.Rule {
	return parser.NewRule().
		Name("header").
		Required().
		Pattern("===").
		Handler(func(ctx *parser.Context) {
			ctx.IgnoreLine()
			ctx.DisableRule()
		})
}

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

func Label(p *page.Page) *parser.Rule {
	return parser.NewRule().
		Name("label").
		Pattern("(?P<indent>\\s*)(?P<label>[^\\s].+?)?(?:\\[(?P<link>.*)\\])?").
		Handler(func(ctx *parser.Context) {
			indent := ctx.Param("indent")
			label := strings.TrimSpace(ctx.Param("label"))
			link := strings.TrimSpace(ctx.Param("link"))

			if len(indent)%4 != 0 {
				ctx.Error("indentation must be in 4 space increments")
				return
			}
			indentLevel := len(indent) / 4

			err := p.AddItem(indentLevel, page.NewItem(link, label))
			if err != nil {
				ctx.Error(err.Error())
				return
			}
			ctx.IgnoreLine()
		})
}
