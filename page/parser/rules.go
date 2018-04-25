package pageparser

import (
	"github.com/g-harel/targetblank/page"
	"github.com/g-harel/targetblank/page/parser/v1"
	"github.com/g-harel/targetblank/parser"
)

// EmptyRule removes lines that are entirely whitespace.
func EmptyRule(p *page.Page) *parser.Rule {
	return parser.NewRule().
		Name("empty").
		Pattern("\\s*").
		Handler(func(ctx *parser.Context) {
			ctx.IgnoreLine()
		})
}

// WhitespaceRule removes empty WhitespaceRule at the end of lines.
func WhitespaceRule(p *page.Page) *parser.Rule {
	return parser.NewRule().
		Name("whitespace").
		Pattern("(?P<content>.+?)\\s+").
		Handler(func(ctx *parser.Context) {
			ctx.ReplaceLine(ctx.Param("content"))
		})
}

// CommentRule removes comments.
func CommentRule(p *page.Page) *parser.Rule {
	return parser.NewRule().
		Name("comment").
		Pattern("(?P<content>[^#]*)(#.*)").
		Handler(func(ctx *parser.Context) {
			ctx.ReplaceLine(ctx.Param("content"))
		})
}

// VersionRule is a required rule which matches with a VersionRule declaration.
// When the VersionRule is found, the corresponding rules are added to the parser.
func VersionRule(p *page.Page) *parser.Rule {
	return parser.NewRule().
		Name("version").
		Required().
		Pattern("version (?P<number>\\d+)").
		Handler(func(ctx *parser.Context) {
			version := ctx.Param("number")
			if version == "1" {
				ctx.AddRules(
					v1.MetadataRule(p),
					v1.HeaderRule(p),
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
