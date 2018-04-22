package pageparser

import (
	"github.com/g-harel/targetblank/page"
	v1rules "github.com/g-harel/targetblank/page/parser/v1"
	"github.com/g-harel/targetblank/parser"
)

// Empty removes lines that are entirely whitespace.
func Empty(p *page.Page) *parser.Rule {
	return parser.NewRule().
		Name("empty").
		Pattern("\\s*").
		Handler(func(ctx *parser.Context) {
			ctx.IgnoreLine()
		})
}

// Whitespace removes empty whitespace at the end of lines.
func Whitespace(p *page.Page) *parser.Rule {
	return parser.NewRule().
		Name("whitespace").
		Pattern("(?P<content>.+?)\\s+").
		Handler(func(ctx *parser.Context) {
			ctx.ReplaceLine(ctx.Param("content"))
		})
}

// Comment removes comments.
func Comment(p *page.Page) *parser.Rule {
	return parser.NewRule().
		Name("comment").
		Pattern("(?P<content>[^#]*)(#.*)").
		Handler(func(ctx *parser.Context) {
			ctx.ReplaceLine(ctx.Param("content"))
		})
}

// Version is a required rule which matches with a version declaration.
// When the version is found, the corresponding rules are added to the parser.
func Version(p *page.Page) *parser.Rule {
	return parser.NewRule().
		Name("version").
		Required().
		Pattern("version (?P<number>\\d+)").
		Handler(func(ctx *parser.Context) {
			version := ctx.Param("number")
			if version == "1" {
				ctx.AddRules(
					v1rules.Metadata(p),
					v1rules.Header(p),
				)
			} else {
				ctx.Error("unsupported version")
				return
			}
			p.SetVersion(version)
			ctx.IgnoreLine()
			ctx.DisableRule()
		})
}
