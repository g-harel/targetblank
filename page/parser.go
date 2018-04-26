package page

import (
	"regexp"

	"github.com/g-harel/targetblank/parser"
)

// NewFromSpec creates a new page from an input specification.
func NewFromSpec(s string) (*Page, *parser.Error) {
	p := New()

	// v1MetadataRule matches with header metadata values.
	v1MetadataRule := parser.Rule{
		Name:    "metadata",
		Pattern: regexp.MustCompile("^(?P<key>[A-Za-z0-9_-]+)\\s*=\\s*(?P<value>.*)\\s*$"),
		Handler: func(ctx *parser.Context) {
			p.AddMeta(ctx.Param("key"), ctx.Param("value"))
			ctx.IgnoreLine()
		},
	}

	// v1GroupRule matches group delimiters.
	// These delimiters indicate a new group should be created.
	v1GroupRule := parser.Rule{
		Name:    "group",
		Pattern: regexp.MustCompile("^---$"),
		Handler: func(ctx *parser.Context) {
			p.AddGroup()
			ctx.IgnoreLine()
			ctx.DisableOtherRule(v1MetadataRule.Name)
		},
	}

	// v1LabelRule matches labelled links.
	// Item are added to the page at the specified depth.
	v1LabelRule := parser.Rule{
		Name:    "label",
		Pattern: regexp.MustCompile("^(?P<indent>\\s*)(?P<label>[^\\s\\[].+?)?(?:\\[(?P<link>.*)\\])?$"),
		Handler: func(ctx *parser.Context) {
			indent := ctx.Param("indent")
			label := ctx.Param("label")
			link := ctx.Param("link")

			if len(indent)%4 != 0 {
				ctx.Error("expected indentation to be in 4 space increments")
				return
			}
			depth := len(indent) / 4

			err := p.AddItem(depth, newItem(link, label))
			if err != nil {
				ctx.Error(err.Error())
				return
			}
			ctx.IgnoreLine()
		},
	}

	// v1HeaderRule is a required rule which matches with the header delimiter.
	// Once the header is found, the remaining rules are added to the parser.
	v1HeaderRule := parser.Rule{
		Name:     "header",
		Required: true,
		Pattern:  regexp.MustCompile("^===$"),
		Handler: func(ctx *parser.Context) {
			ctx.IgnoreLine()
			ctx.DisableRule()
			ctx.AddRules(
				v1GroupRule,
				v1LabelRule,
			)
		},
	}

	// emptyRule removes lines that are entirely whitespace.
	emptyRule := parser.Rule{
		Name:    "empty",
		Pattern: regexp.MustCompile("^\\s*$"),
		Handler: func(ctx *parser.Context) {
			ctx.IgnoreLine()
		},
	}

	// whitespaceRule removes empty whitespace at the end of lines.
	whitespaceRule := parser.Rule{
		Name:    "whitespace",
		Pattern: regexp.MustCompile("^(?P<content>.+?)\\s+$"),
		Handler: func(ctx *parser.Context) {
			ctx.ReplaceLine(ctx.Param("content"))
		},
	}

	// commentRule removes comments.
	commentRule := parser.Rule{
		Name:    "comment",
		Pattern: regexp.MustCompile("^(?P<content>[^#]*)(#.*)$"),
		Handler: func(ctx *parser.Context) {
			ctx.ReplaceLine(ctx.Param("content"))
		},
	}

	// versionRule is a required rule which matches with a version declaration.
	// When the version is found, the corresponding rules are added to the parser.
	versionRule := parser.Rule{
		Name:     "version",
		Required: true,
		Pattern:  regexp.MustCompile("^version (?P<number>\\d+)$"),
		Handler: func(ctx *parser.Context) {
			version := ctx.Param("number")
			if version == "1" {
				ctx.AddRules(
					v1MetadataRule,
					v1HeaderRule,
				)
			} else {
				ctx.Error("unsupported version")
				return
			}
			p.SetVersion(version)
			ctx.IgnoreLine()
			ctx.DisableRule()
		},
	}

	ps := parser.New()
	ps.Add(
		emptyRule,
		whitespaceRule,
		commentRule,
		versionRule,
	)
	return p, ps.Parse(s)
}
