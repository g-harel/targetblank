package pageparser

import (
	"github.com/g-harel/targetblank/page"
	v1rules "github.com/g-harel/targetblank/page/parser/v1"
	"github.com/g-harel/targetblank/parser"
)

func Empty(p *page.Page) *parser.Rule {
	return parser.NewRule().
		Name("empty").
		Pattern("\\s*").
		Handler(func(ctx *parser.Context) {
			ctx.IgnoreLine()
		})
}

func Whitespace(p *page.Page) *parser.Rule {
	return parser.NewRule().
		Name("whitespace").
		Pattern("(?P<content>.+?)\\s+").
		Handler(func(ctx *parser.Context) {
			ctx.ReplaceLine(ctx.Param("content"))
		})
}

func Comment(p *page.Page) *parser.Rule {
	return parser.NewRule().
		Name("comment").
		Pattern("(?P<content>[^#]*)(#.*)").
		Handler(func(ctx *parser.Context) {
			ctx.ReplaceLine(ctx.Param("content"))
		})
}

func Version(p *page.Page) *parser.Rule {
	return parser.NewRule().
		Name("version").
		Required().
		Pattern("version (?P<number>\\d+)").
		Handler(func(ctx *parser.Context) {
			number := ctx.Param("number")
			if number == "1" {
				ctx.AddRules(
					v1rules.Metadata(p),
					v1rules.Header(p),
					v1rules.Group(p),
					v1rules.Label(p),
				)
			} else {
				ctx.Error("unsupported version")
				return
			}
			p.SetVersion(number)
			ctx.IgnoreLine()
			ctx.DisableRule()
		})
}
